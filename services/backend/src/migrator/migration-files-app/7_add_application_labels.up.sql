CREATE TABLE "application_label_colors" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" VARCHAR(255) NOT NULL UNIQUE CHECK (name <> ''),
  
  -- check: either 3 or 6 digits, or 3 or 6 digits with alpha channel (= 8 or 5 digits)
  "border_color_light" VARCHAR(9) NOT NULL CHECK (border_color_light ~ '^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3}|[0-9a-fA-F]{8}|[0-9a-fA-F]{5})$'), 
  "bg_color_light" VARCHAR(9) NOT NULL CHECK (bg_color_light ~ '^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3}|[0-9a-fA-F]{8}|[0-9a-fA-F]{5})$'), 
  "color_light" VARCHAR(9) NOT NULL CHECK (color_light ~ '^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3}|[0-9a-fA-F]{8}|[0-9a-fA-F]{5})$'), 

  "border_color_dark" VARCHAR(9) NOT NULL CHECK (border_color_dark ~ '^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3}|[0-9a-fA-F]{8}|[0-9a-fA-F]{5})$'), 
  "bg_color_dark" VARCHAR(9) NOT NULL CHECK (bg_color_dark ~ '^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3}|[0-9a-fA-F]{8}|[0-9a-fA-F]{5})$'), 
  "color_dark" VARCHAR(9) NOT NULL CHECK (color_dark ~ '^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3}|[0-9a-fA-F]{8}|[0-9a-fA-F]{5})$'), 

  UNIQUE (bg_color_dark, color_dark, border_color_dark),
  UNIQUE (bg_color_light, color_light, border_color_light)
);

INSERT INTO "application_label_colors" ("name", "bg_color_dark", "color_dark", "border_color_dark", "bg_color_light", "color_light", "border_color_light") VALUES
('red', '#ff00004d', '#ffcccc','#ffcccc', '#ffebee', '#c62828', '#c62828'),
('green', '#00ff004d', '#ccffcc','#ccffcc', '#e8f5e8', '#2e7d32', '#2e7d32'),
('blue', '#0000ff4d', '#ccccff','#ccccff', '#e3f2fd', '#1565c0', '#1565c0'),
('yellow', '#ffff004d', '#ffffcc','#ffffcc', '#fffde7', '#f57f17', '#f57f17'),
('purple', '#8000804d', '#e6ccff','#e6ccff', '#f3e5f5', '#7b1fa2', '#7b1fa2'),
('orange', '#c3641ff2', '#e2c3af','#e2c3af', '#fff3e0', '#ef6c00', '#ef6c00'),
('pink', '#ba4e78b5', '#dec7d4','#dec7d4', '#ffe4f5', '#c2185b', '#c2185b'),
('gray', '#8080804d', '#e6e6e6','#e6e6e6', '#f5f5f5', '#424242', '#424242');

CREATE TABLE "application_labels" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" VARCHAR(255) NOT NULL UNIQUE CHECK (name <> ''),
  color_id UUID NOT NULL REFERENCES "application_label_colors" ("id"),
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  UNIQUE (name, color_id)
);

INSERT INTO "application_labels" ("name", "color_id") VALUES
('red', (SELECT id FROM application_label_colors WHERE name = 'red')),
('green', (SELECT id FROM application_label_colors WHERE name = 'green')),
('blue', (SELECT id FROM application_label_colors WHERE name = 'blue')),
('yellow', (SELECT id FROM application_label_colors WHERE name = 'yellow')),
('purple', (SELECT id FROM application_label_colors WHERE name = 'purple')),
('orange', (SELECT id FROM application_label_colors WHERE name = 'orange')),
('pink', (SELECT id FROM application_label_colors WHERE name = 'pink')),
('gray', (SELECT id FROM application_label_colors WHERE name = 'gray'));


CREATE VIEW application_labels_with_color AS (
   SELECT al.id,
    al.name,
	json_build_object(
        'id', alc.id, 
		    'name', alc.name,
        'bg_color_light', alc.bg_color_light,
        'color_light', alc.color_light,
        'bg_color_dark', alc.bg_color_dark,
        'color_dark', alc.color_dark
    ) AS style
   FROM application_labels al
     JOIN application_label_colors alc ON al.color_id = alc.id
  GROUP BY al.id, alc.id
);

CREATE TABLE application__application_labels (
	"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	"application_id" uuid NOT NULL,
	"application_label_id" uuid NOT NULL,
	UNIQUE ("application_id", "application_label_id")
);

DROP VIEW IF EXISTS "applications_view";
CREATE VIEW "applications_view" AS (
SELECT 
    app.id,
    app.class_level AS class_level,
    COALESCE(
      jsonb_agg(
        jsonb_build_object(
          'id', application_labels.id,
          'name', application_labels.name,
          'color', jsonb_build_object(
            'id', application_label_colors.id,
            'name', application_label_colors.name,
            'border_color_light', application_label_colors."border_color_light",
            'bg_color_light', application_label_colors."bg_color_light",
            'color_light', application_label_colors."color_light",
            'border_color_dark', application_label_colors."border_color_dark",
            'bg_color_dark', application_label_colors."bg_color_dark",
            'color_dark', application_label_colors."color_dark"
          )
        )
      ) FILTER (WHERE application_labels.id IS NOT NULL),
      '[]'::jsonb
    ) AS "labels",
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
     LEFT JOIN users assigned_user ON app.assigned_user_id = assigned_user.id
     JOIN applicants applicant ON app.applicant_id = applicant.id
     JOIN applicant_permanent_address applicant_permanent_address ON applicant.address_id = applicant_permanent_address.id
      JOIN applicants_contact_data applicant_contact_data ON applicant.contact_id = applicant_contact_data.id
     JOIN school_degrees school_degree ON app.school_degree_id = school_degree.id
     JOIN schools ON school_degree.school_id = schools.id
     JOIN school_types ON schools.school_type_id = school_types.id
     LEFT JOIN application__application_labels rel_application_labels ON rel_application_labels.application_id = app.id
     LEFT JOIN application_labels ON application_labels.id = rel_application_labels.application_label_id
     LEFT JOIN application_label_colors ON application_labels.color_id = application_label_colors.id
	GROUP BY app.id, app_status.id,applicant.id, applicant_permanent_address.id, applicant_contact_data.id, school_degree.id, school_types.id, schools.id, assigned_user.id
  ORDER BY app.created
);