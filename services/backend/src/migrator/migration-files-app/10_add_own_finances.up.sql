CREATE TABLE applicant_own_income_and_finances(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  application_revision_id UUID NOT NULL,
  
  bewilligungszeitraum_start DATE NOT NULL,
  bewilligungszeitraum_end DATE NOT NULL
);