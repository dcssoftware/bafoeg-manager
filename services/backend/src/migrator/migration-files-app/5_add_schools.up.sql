CREATE TABLE "school_types" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" varchar(255),
  "identifier" varchar(255) UNIQUE
);

INSERT INTO "school_types" (name, identifier) VALUES
  ('Gymnasium', 'high_school'),
  ('Realschule', 'middle_school'),
  ('Hauptschule', 'lower_middle_school'),
  ('Gesamtschule', 'comprehensive_school'),
  ('Abendschule', 'evening_school'),
  ('Berufsaufbauschule', 'vocational_school'),
  ('Abendrealschule', 'evening_middle_school'),
  ('Abendgymnaium', 'evening_high_school'),
  ('Kolleg', 'college'),
  ('Sonstige', 'other_school');

CREATE TABLE "schools" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" varchar,
  "school_type_id" UUID NOT NULL REFERENCES "school_types" ("id"),
  "street" varchar,
  "house_number" varchar,
  "zip_code" varchar,
  "city" varchar,
  "country" varchar
);

CREATE TABLE "school_degrees" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "school_id" UUID NOT NULL REFERENCES "schools" ("id"),
  "name" varchar NOT NULL CHECK (name <> ''),
  "fos_berufsabschluss_required" boolean NOT NULL,
  "bos_berufsqualifizierender_abschluss" boolean NOT NULL,
  "fachschule_berufsabschluss_required" boolean NOT NULL
);

 INSERT INTO "schools" (name, school_type_id, street, house_number, zip_code, city, country) VALUES (
    'Test Schule (example school)', 
    (SELECT id FROM school_types WHERE identifier = 'other_school'), 
    'Teststraße', 
    '1a', 
    '12345', 
    'Teststadt', 
    'Testland'
  );

  INSERT INTO "school_degrees" (school_id, name, fos_berufsabschluss_required, bos_berufsqualifizierender_abschluss, fachschule_berufsabschluss_required) VALUES (
    (SELECT id FROM schools WHERE name = 'Test Schule (example school)'), 
    'Test Abschluss (example degree)',
    false,
    true,
    false
  );
  
  CREATE VIEW school_overview AS (
		SELECT 
			schools.id, 
			schools.name, 
			schools.street, 
			schools.house_number, 
			schools.city, 
			schools.zip_code, 
			schools.country,
			school_types.name AS "school_type_name",
			school_types.identifier AS "school_type_identifier"
		FROM schools
		
		INNER JOIN school_types
			ON schools.school_type_id = school_types.id 
);