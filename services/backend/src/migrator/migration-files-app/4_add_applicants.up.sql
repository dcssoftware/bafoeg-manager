CREATE TABLE "applicants" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "contact_id" UUID,
  "address_id" UUID,
  "bank_account_id" UUID
);

CREATE TABLE "applicant_permanent_address" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "street" varchar,
  "house_number" varchar,
  "zip_code" varchar,
  "city" varchar,
  "country" varchar
);

CREATE TABLE "applicants_contact_data" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "phone" varchar,
  "email" varchar
);

CREATE TABLE "applicants_bank_data" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "iban" varchar,
  "bic" varchar DEFAULT NULL,
  "bank_name" varchar DEFAULT NULL,
  "account_holder" varchar DEFAULT NULL
);

ALTER TABLE "applicants" ADD FOREIGN KEY ("contact_id") REFERENCES "applicants_contact_data" ("id");
ALTER TABLE "applicants" ADD FOREIGN KEY ("address_id") REFERENCES "applicant_permanent_address" ("id");
ALTER TABLE "applicants" ADD FOREIGN KEY ("bank_account_id") REFERENCES "applicants_bank_data" ("id");

CREATE VIEW "applicants_with_address_and_contact_data" AS (
SELECT 
	applicant.id AS "id",
	json_build_object(
		'id', applicant.id,
		'firstname', applicant.firstname,
		'lastname', applicant.lastname,
		'address', json_build_object(
			'street', address.street,
			'house_number', address."house_number",
			'zip_code', address."zip_code",
			'city', address.city,
			'country', address.country
		),
		'contact_data', json_build_object(
			'email', contact.email,
			'phone', contact.phone
		)
	) AS "applicant"

	FROM applicants applicant

	INNER JOIN applicant_permanent_address address
		ON applicant.address_id = address.id

	INNER JOIN applicants_contact_data contact
		ON applicant.contact_id = contact.id
);