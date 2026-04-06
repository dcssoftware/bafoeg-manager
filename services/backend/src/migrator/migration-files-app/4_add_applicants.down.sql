ALTER TABLE "applicants" DROP CONSTRAINT "applicants_contact_id_fkey";
ALTER TABLE "applicants" DROP CONSTRAINT "applicants_address_id_fkey";
ALTER TABLE "applicants" DROP CONSTRAINT "applicants_bank_account_id_fkey";

DROP TABLE IF EXISTS "applicants";
DROP TABLE IF EXISTS "applicants_bank_data";
DROP TABLE IF EXISTS "applicants_contact_data";
DROP TABLE IF EXISTS "applicant_permanent_address";