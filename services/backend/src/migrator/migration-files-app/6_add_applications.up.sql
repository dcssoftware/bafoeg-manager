CREATE TYPE application_status_enum AS ENUM ();
ALTER TYPE application_status_enum ADD VALUE 'IN_PROGRESS'; 
ALTER TYPE application_status_enum ADD VALUE 'RESPONSE_AWAITED'; 
ALTER TYPE application_status_enum ADD VALUE 'DENIED'; 
ALTER TYPE application_status_enum ADD VALUE 'APPROVED';

CREATE TABLE "application_status" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "identifier" application_status_enum NOT NULL UNIQUE,
  "name" varchar NOT NULL
);

INSERT INTO "application_status" (identifier, name) VALUES
  ('IN_PROGRESS', 'In Progress'),
  ('RESPONSE_AWAITED', 'Response Awaited'),
  ('DENIED', 'Denied'),
  ('APPROVED', 'Approved');

-- Function to calculate application_validity_starts
CREATE OR REPLACE FUNCTION calculate_application_default_next_available_validity_starts() 
RETURNS timestamp without time zone AS $$
BEGIN
  RETURN CASE 
    WHEN (CURRENT_DATE <= make_date(EXTRACT(YEAR FROM CURRENT_DATE)::INTEGER, 8, 1))
    THEN make_date(EXTRACT(YEAR FROM CURRENT_DATE)::INTEGER, 9, 1)
    ELSE make_date(EXTRACT(YEAR FROM CURRENT_DATE)::INTEGER + 1, 9, 1)
  END;
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to calculate application_validity_ends
CREATE OR REPLACE FUNCTION calculate_application_default_next_available_validity_ends() 
RETURNS timestamp without time zone AS $$
BEGIN
  RETURN CASE 
    WHEN (CURRENT_DATE <= make_date(EXTRACT(YEAR FROM CURRENT_DATE)::INTEGER, 7, 31))
    THEN make_date(EXTRACT(YEAR FROM CURRENT_DATE)::INTEGER, 7, 31)
    ELSE make_date(EXTRACT(YEAR FROM CURRENT_DATE)::INTEGER + 1, 7, 31)
  END;
END;
$$ LANGUAGE plpgsql STABLE;

CREATE TABLE "applications" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "applicant_id" uuid NOT NULL,
  
  "class_level" varchar NOT NULL CHECK (class_level <> ''),
  "status" uuid NOT NULL,

  "assigned_user_id" uuid,
  "school_degree_id" uuid NOT NULL,

  "application_validity_starts" timestamp without time zone NOT NULL DEFAULT calculate_application_default_next_available_validity_starts(),
  "application_validity_ends" timestamp without time zone NOT NULL DEFAULT calculate_application_default_next_available_validity_ends(),
  "created" timestamp without time zone NOT NULL DEFAULT now(),
  "updated" timestamp without time zone NOT NULL DEFAULT now()
);



-- Add constraint to ensure application_validity_starts is before application_validity_ends
ALTER TABLE "applications" ADD CONSTRAINT "application_validity_order_check" 
CHECK (application_validity_starts < application_validity_ends);

-- Add constraint to ensure application_validity_ends is at least 1 month after application_validity_starts
ALTER TABLE "applications" ADD CONSTRAINT "application_validity_minimum_duration_check" 
CHECK (application_validity_ends >= application_validity_starts + INTERVAL '1 month');

-- Add constraint to ensure application_validity_ends is at most 1 year and 3 months after application_validity_starts
ALTER TABLE "applications" ADD CONSTRAINT "application_validity_maximum_duration_check" 
CHECK (application_validity_ends <= application_validity_starts + INTERVAL '1 year 3 months');

-- Add btree_gist extension for exclusion constraints with btree types
CREATE EXTENSION IF NOT EXISTS btree_gist;

-- Create immutable wrapper function for tstzrange
CREATE OR REPLACE FUNCTION create_application_timerange(start_time timestamp without time zone, end_time timestamp without time zone) 
RETURNS tstzrange AS $$
BEGIN
  RETURN tstzrange(start_time, end_time);
END;
$$ LANGUAGE plpgsql IMMUTABLE;

-- Add exclusion constraint to prevent overlapping application periods for the same applicant
ALTER TABLE "applications" ADD CONSTRAINT "no_overlapping_applications_for_applicant"
EXCLUDE USING gist (applicant_id WITH =, 
                   create_application_timerange(application_validity_starts, application_validity_ends) WITH &&);

CREATE TABLE "application_revisions" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "message_header" VARCHAR(255) NOT NULL,
  "message_description" varchar NOT NULL,
  "application_id" uuid NOT NULL,
  "trainings_address_id" uuid,
  "created" timestamp without time zone NOT NULL DEFAULT now()
);  

CREATE TABLE "applicant_training_address" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "street" varchar,
  "house_number" varchar,
  "zip_code" varchar,
  "city" varchar,
  "country" varchar
);

ALTER TABLE "applications" ADD FOREIGN KEY ("status") REFERENCES "application_status" ("id");
ALTER TABLE "applications" ADD FOREIGN KEY ("applicant_id") REFERENCES "applicants" ("id");
ALTER TABLE "applications" ADD FOREIGN KEY ("assigned_user_id") REFERENCES "users" ("id");
ALTER TABLE "applications" ADD FOREIGN KEY ("school_degree_id") REFERENCES "school_degrees" ("id");
ALTER TABLE "application_revisions" ADD FOREIGN KEY ("application_id") REFERENCES "applications" ("id");
ALTER TABLE "application_revisions" ADD FOREIGN KEY ("trainings_address_id") REFERENCES "applicant_training_address" ("id");

CREATE VIEW "applications_overview" AS (
  SELECT 
			app.id,
      app.class_level AS "class_level",
      assigned_user.id AS "assigned_user_id",
      applicant.id AS applicant_id,
      app_status.identifier AS status_identifier,

      json_build_object(
        'id', applicant.id,
        'firstname', applicant.firstname,
        'lastname', applicant.lastname
      ) AS "applicant",

      json_build_object(
        'id', assigned_user.id,
        'username', assigned_user.username,
        'display_name', assigned_user.display_name
      ) AS "assigned_user",

      json_build_object(
        'id', schools.id,
        'name', schools.name,
        'address', json_build_object(
          'street', schools.street,
          'house_number', schools."house_number", 
          'zip_code',schools."zip_code",
          'city', schools.city,
          'country',schools.country
        ),
        'type', json_build_object(
          'id', school_types.id,
          'name', school_types.name,
          'identifier', school_types.identifier
        ),
        'degree', json_build_object(
          'id',school_degree.id,
          'name',school_degree.name,
          'fos_berufsabschluss_required',school_degree.fos_berufsabschluss_required,
          'bos_berufsqualifizierender_abschluss',school_degree.bos_berufsqualifizierender_abschluss,
          'fachschule_berufsabschluss_required',school_degree.fachschule_berufsabschluss_required
        ) 
      )AS "school",

      json_build_object(
      'id', app_status.id,
        'name', app_status.name,
        'identifier', app_status.identifier
      ) AS "status",

			app.created AS "application_created",
      app.updated AS "application_updated"

		FROM public.applications app

    INNER JOIN application_status app_status
      ON app.status = app_status.id
		LEFT JOIN users assigned_user
			ON app.assigned_user_id = assigned_user.id
		INNER JOIN applicants applicant
			ON app.applicant_id = applicant.id
    INNER JOIN school_degrees school_degree
      ON app.school_degree_id = school_degree.id
		INNER JOIN schools 
			ON school_degree.school_id = schools.id
		INNER JOIN school_types
			ON schools.school_type_id = school_types.id
    
    GROUP BY app.id, app_status.id,applicant.id, school_degree.id, school_types.id, schools.id, assigned_user.id
    ORDER BY app.created
);

CREATE VIEW "applications_view" AS (
SELECT 
    app.id,
    app.class_level AS class_level,
    json_build_object(
      'id', applicant.id,
      'firstname', applicant.firstname,
      'lastname', applicant.lastname,
      'address', json_build_object(
        'street', applicant_permanent_address.street,
        'house_number',applicant_permanent_address."house_number", 
        'zip_code',applicant_permanent_address."zip_code",
        'city', applicant_permanent_address.city,
        'country',applicant_permanent_address.country
      ),
      'contact_data', json_build_object(
        'email',applicant_contact_data.email,
        'phone',applicant_contact_data.phone
      )
    ) AS "applicant",
    json_build_object(
      'id', schools.id,
      'name', schools.name,
      'address', json_build_object(
        'street', schools.street,
        'house_number', schools."house_number", 
        'zip_code',schools."zip_code",
        'city', schools.city,
        'country',schools.country
      ),
      'type', json_build_object(
        'id', school_types.id,
        'name', school_types.name,
        'identifier', school_types.identifier
      ),
      'degree', json_build_object(
        'id',school_degree.id,
        'name',school_degree.name,
        'fos_berufsabschluss_required',school_degree.fos_berufsabschluss_required,
        'bos_berufsqualifizierender_abschluss',school_degree.bos_berufsqualifizierender_abschluss,
        'fachschule_berufsabschluss_required',school_degree.fachschule_berufsabschluss_required
      ) 
    )AS "school",
    CASE
    WHEN assigned_user.id IS NOT NULL THEN
        json_build_object(
            'id', assigned_user.id,
            'username', assigned_user.username,
            'display_name', assigned_user.display_name
        )
      ELSE NULL
    END AS assigned_user,
    json_build_object(
	  'id', app_status.id,
      'name', app_status.name,
      'identifier', app_status.identifier
    ) AS "status",
    app.created AS application_created,
    app.updated AS application_updated
   FROM applications app
     JOIN application_status app_status ON app.status = app_status.id
     JOIN users assigned_user ON app.assigned_user_id = assigned_user.id
     JOIN applicants applicant ON app.applicant_id = applicant.id
     JOIN applicant_permanent_address applicant_permanent_address ON applicant.address_id = applicant_permanent_address.id
      JOIN applicants_contact_data applicant_contact_data ON applicant.contact_id = applicant_contact_data.id
     JOIN school_degrees school_degree ON app.school_degree_id = school_degree.id
     JOIN schools ON school_degree.school_id = schools.id
     JOIN school_types ON schools.school_type_id = school_types.id
	GROUP BY app.id, app_status.id,applicant.id, applicant_permanent_address.id, applicant_contact_data.id, school_degree.id, school_types.id, schools.id, assigned_user.id
  ORDER BY app.created
);

CREATE VIEW "school_applicants_view" AS (
SELECT
    school.id AS "school_id",
    app.class_level AS "application_class_level",
    application_status.identifier AS application_status_identifier,
    app.application_validity_starts AS "application_validity_starts",
    app.application_validity_ends AS "application_validity_ends",
    applicants.id AS "applicants_id",
    applicants.firstname AS "applicants_firstname",
    applicants.lastname AS "applicants_lastname",
    applicant_permanent_address."zip_code" AS "applicants_address_zip_code",
    applicant_permanent_address.city AS "applicants_address_city",
    applicant_permanent_address.country AS "applicants_address_country",
    school_degree.id AS "school_degree_id",
    school_degree.name AS "school_degree_name"
FROM applications app
INNER JOIN applicants applicants 
    ON app.applicant_id = applicants.id
INNER JOIN applicant_permanent_address applicant_permanent_address
    ON applicants.address_id = applicant_permanent_address.id
INNER JOIN school_degrees school_degree
    ON app.school_degree_id = school_degree.id
INNER JOIN schools school
    ON school_degree.school_id = school.id
JOIN application_status 
    ON app.status = application_status.id
WHERE application_status.identifier != 'DENIED'
ORDER BY applicants.id, app.updated DESC
);

-- Create trigger function to update application timestamp
CREATE OR REPLACE FUNCTION update_application_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  -- For direct updates to applications table
  IF TG_TABLE_NAME = 'applications' THEN
    NEW.updated = now();
    RETURN NEW;
  -- For operations on application_revisions table
  ELSIF TG_TABLE_NAME = 'application_revisions' THEN
    UPDATE applications
    SET updated = now()
    WHERE id = NEW.application_id;
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for applications table
CREATE TRIGGER application_updated
BEFORE UPDATE ON applications
FOR EACH ROW
EXECUTE FUNCTION update_application_timestamp();

-- Create triggers for application_revisions table
CREATE TRIGGER application_updated
AFTER UPDATE ON application_revisions
FOR EACH ROW
EXECUTE FUNCTION update_application_timestamp();

-- Also trigger when new revisions are created
CREATE TRIGGER application_revision_inserted
AFTER INSERT ON application_revisions
FOR EACH ROW
EXECUTE FUNCTION update_application_timestamp();
