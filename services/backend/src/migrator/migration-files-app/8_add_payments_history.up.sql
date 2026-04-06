CREATE TYPE "payment_direction" AS ENUM ('outgoing', 'incoming');

CREATE TABLE "payment_status" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "identifier" VARCHAR(255) NOT NULL UNIQUE CHECK (identifier <> '')
);

INSERT INTO "payment_status" ("identifier") VALUES
('pending'),
('completed'),
('failed'),
('refunded');

CREATE TABLE "payments" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  "applicant_id" uuid NOT NULL,
  "application_id" uuid,
  "amount" NUMERIC(10, 2) NOT NULL CHECK (amount > 0),
  "executed" timestamp without time zone NOT NULL DEFAULT now(),
  "status_id" UUID NOT NULL,
  "description" VARCHAR(255) NOT NULL CHECK (description <> ''),

  "iban" VARCHAR(34) NOT NULL CHECK (iban ~ '^[A-Z]{2}[0-9]{2}[A-Z0-9]{15,30}$'),
  "bic" VARCHAR(11) NOT NULL CHECK (bic ~ '^[A-Z]{6}[A-Z0-9]{0,5}$'),
  "direction" "payment_direction" NOT NULL,

  "created" timestamp without time zone NOT NULL DEFAULT now()
);

ALTER TABLE "payments" ADD FOREIGN KEY ("applicant_id") REFERENCES "applicants" ("id");
ALTER TABLE "payments" ADD FOREIGN KEY ("application_id") REFERENCES "applications" ("id");
ALTER TABLE "payments" ADD FOREIGN KEY ("status_id") REFERENCES "payment_status" ("id");