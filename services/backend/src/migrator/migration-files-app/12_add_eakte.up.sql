CREATE TYPE eakte_akte_vertraulich_enum AS ENUM ();
ALTER TYPE eakte_akte_vertraulich_enum ADD VALUE 'VERTRAULICH'; 
ALTER TYPE eakte_akte_vertraulich_enum ADD VALUE 'INTERN'; 
ALTER TYPE eakte_akte_vertraulich_enum ADD VALUE 'OFFEN'; 

CREATE TABLE eakte_import_akte_type (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR NOT NULL CHECK (name <> ''),
  identifier VARCHAR(255) UNIQUE NOT NULL CHECK (name <> '')
);

INSERT INTO eakte_import_akte_type (name, identifier) VALUES 
('Antrag', 'ANTRAG'),
('Sachakte', 'SACHAKTE');

CREATE TABLE eakte_akte_source (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR NOT NULL CHECK (name <> ''),
  identifier VARCHAR(255) UNIQUE NOT NULL CHECK (name <> '')
);

INSERT INTO eakte_akte_source (name, identifier) VALUES 
('Sharepoint', 'SHAREPOINT'),
('E-Mail', 'EMAIL');

CREATE TABLE eakte_import_akte (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  aktenzeichen VARCHAR NOT NULL CHECK (aktenzeichen <> ''),
  typ UUID NOT NULL REFERENCES eakte_import_akte_type(id),
  vertraulichkeit eakte_akte_vertraulich_enum NOT NULL,
  created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE eakte_import_vorgang (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  akte_id UUID NOT NULL REFERENCES eakte_import_akte(id),
  vorgangszeichen VARCHAR NOT NULL CHECK (vorgangszeichen <> ''),
  created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE eakte_import_dokument (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  source UUID NOT NULL REFERENCES eakte_akte_source(id),
  vorgang_id UUID NOT NULL REFERENCES eakte_import_vorgang(id),
  source_xdomea_file BOOLEAN NOT NULL DEFAULT FALSE,
  file_id UUID NOT NULL REFERENCES files(id),
  created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE application_eakte_mapping (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  application_id UUID NOT NULL REFERENCES applications(id),
  eakte_akte_id UUID NOT NULL REFERENCES eakte_import_akte(id),
  UNIQUE (application_id, eakte_akte_id)
);

CREATE VIEW eakte_overview AS (
 SELECT eakte_import_akte.id,
    eakte_import_akte.aktenzeichen,
    json_build_object('id', eakte_import_akte_type.id, 'name', eakte_import_akte_type.name, 'identifier', eakte_import_akte_type.identifier) AS typ,
    eakte_import_akte.vertraulichkeit,
    eakte_import_akte.created
   FROM eakte_import_akte
     LEFT JOIN eakte_import_akte_type ON eakte_import_akte.typ = eakte_import_akte_type.id
);

CREATE VIEW eakte_files_overview AS (
   SELECT 
  eakte_import_dokument.id, 
  eakte_import_dokument.source_xdomea_file,
  json_build_object(
	'id', eakte_akte_source.id,
	'identifier', eakte_akte_source.identifier,
	'name', eakte_akte_source.name
  ) AS "source",
  json_build_object(
    'id', files.id,
    'name', files.file_name,
    'file_type', files.file_type,
    'file_size', files.file_size,
    'created', to_char(files.created AT TIME ZONE 'Europe/Berlin'::text, 'YYYY-MM-DD"T"HH24:MI:SS.USOF')
  ) AS files,
  json_build_object(
    'id', eakte_import_vorgang.id,
    'vorgangszeichen', eakte_import_vorgang.vorgangszeichen
  ) AS vorgang,
  json_build_object(
    'id', eakte_import_akte.id,
    'aktenzeichen', eakte_import_akte.aktenzeichen,
    'vertraulichkeit', eakte_import_akte.vertraulichkeit
  ) AS akte,
  eakte_import_dokument.created

 FROM eakte_import_dokument
 JOIN files ON eakte_import_dokument.file_id = files.id
 JOIN eakte_import_vorgang ON eakte_import_dokument.vorgang_id = eakte_import_vorgang.id
 JOIN eakte_import_akte ON eakte_import_vorgang.akte_id = eakte_import_akte.id
 JOIN eakte_akte_source ON eakte_akte_source.id = eakte_import_dokument.source
);

CREATE VIEW eakte_application_mappings_overview AS (
  SELECT 
    application_eakte_mapping.id AS id,
    application_eakte_mapping.application_id AS application_id,
    application_eakte_mapping.eakte_akte_id AS eakte_akte_id,
    json_build_object(
      'id', applications.id,
      'class', applications.class_level,
      'status', json_build_object(
        'id', application_status.id,
        'identifier', application_status.identifier,
        'name', application_status.name
      )

    ) AS application,
    json_build_object(
      'id', applicants.id,
      'firstname', applicants.firstname,
      'lastname', applicants.lastname
    ) AS applicant,
    json_build_object(
      'id', school_degrees.id,
      'name', school_degrees.name
    ) AS school_degree,
    json_build_object(
      'id', schools.id,
      'name', schools.name
    ) AS school,
    json_build_object(
      'id', eakte_import_akte.id,
      'aktenzeichen', eakte_import_akte.aktenzeichen
    ) AS eakte_akte
    
  FROM application_eakte_mapping
  JOIN applications ON application_eakte_mapping.application_id = applications.id
  JOIN application_status ON applications.status = application_status.id
  JOIN applicants ON applications.applicant_id = applicants.id
  JOIN school_degrees ON applications.school_degree_id = school_degrees.id
  JOIN schools ON school_degrees.school_id = schools.id
  JOIN eakte_import_akte ON application_eakte_mapping.eakte_akte_id = eakte_import_akte.id
);

CREATE VIEW application_files_view AS (
   SELECT 'APPLICATION_DOCUMENT'::text AS source,
    application_files.id,
    application_files.application_id,
    NULL::jsonb AS eakte,
    jsonb_build_object(
      'file_id', files.id, 
      'name', files.file_name, 
      'type', files.file_type, 
      'size', files.file_size, 
      'created', files.created::timestamp with time zone
      ) AS file
   FROM application_files
     JOIN files ON application_files.file_id = files.id
UNION
 SELECT 'EAKTE_DOCUMENT'::text AS source,
    eakte_import_dokument.id,
    COALESCE(application_eakte_mapping.application_id, NULL::uuid) AS application_id,
    jsonb_build_object(
      'akte_id', eakte_import_akte.id, 
      'vertraulichkeit', eakte_import_akte.vertraulichkeit, 
      'created', eakte_import_akte.created::timestamp with time zone
      ) AS eakte,
    jsonb_build_object(
      'file_id', files.id, 
      'name', files.file_name, 
      'type', files.file_type, 
      'size', files.file_size, 
      'created', files.created::timestamp with time zone
      ) AS file
   FROM eakte_import_dokument
     LEFT JOIN eakte_import_vorgang ON eakte_import_dokument.vorgang_id = eakte_import_vorgang.id
     LEFT JOIN eakte_import_akte ON eakte_import_vorgang.akte_id = eakte_import_akte.id
     LEFT JOIN files ON eakte_import_dokument.file_id = files.id
     LEFT JOIN application_eakte_mapping ON application_eakte_mapping.eakte_akte_id = eakte_import_akte.id
  WHERE eakte_import_dokument.source_xdomea_file = false
);