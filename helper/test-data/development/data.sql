INSERT INTO public.applicants_contact_data (email, phone) VALUES
    ('max.mustermann@example.com', '0123456789') RETURNING id;

INSERT INTO public.applicant_permanent_address (street, house_number, zip_code, city, country) VALUES
    ('Musterstraße', '1', '12345', 'Musterstadt', 'Deutschland') RETURNING id;

INSERT INTO public.applicants_bank_data (iban, bic, bank_name, account_holder) VALUES
    ('DE89370400440532013000', 'COBADEFFXXX', 'Commerzbank', 'Max Mustermann') RETURNING id;

INSERT INTO public.applicants (firstname, lastname, contact_id, address_id, bank_account_id) VALUES
(
  'Max', 
  'Mustermann', 
  (
    SELECT id FROM public.applicants_contact_data WHERE email = 'max.mustermann@example.com'
  ), 
  (
    SELECT id FROM public.applicant_permanent_address WHERE street = 'Musterstraße'
  ), 
  (
    SELECT id FROM public.applicants_bank_data WHERE iban = 'DE89370400440532013000'
  )
);

INSERT INTO public.applications (applicant_id, class_level, status, assigned_user_id, school_degree_id, application_validity_starts, application_validity_ends) 
VALUES(
  (
    SELECT id FROM public.applicants WHERE firstname = 'Max' AND lastname = 'Mustermann'
  ),
  '3. Semester',
  (
    SELECT id FROM public.application_status WHERE identifier = 'IN_PROGRESS'
  ),
  (
    SELECT id FROM public.users LIMIT 1
  ),
  (
    SELECT id FROM public.school_degrees WHERE name = 'Test Abschluss (example degree)'
  ),
  '2025-09-01 00:00:00',
  '2026-08-31 00:00:00'
);


INSERT INTO applicant_training_address 
(street, house_number, zip_code, city, country) 
VALUES 
(
'Musterstraße',
'123a',
'80633',
'München',
'Deutschland'
);

INSERT INTO application_revisions 
(message_header, message_description, application_id, trainings_address_id)
VALUES
(
'test',
'test',
(SELECT id FROM applications LIMIT 1),
(SELECT id FROM applicant_training_address LIMIT 1)
);

INSERT INTO user_roles
(user_id, role_id)
VALUES
(
(SELECT id FROM users LIMIT 1),
(SELECT id FROM roles WHERE identifier = 'admin')
);

-- INSERT INTO payments
-- (applicant_id, application_id, amount, executed, status_id, 
-- 	description, iban, bic, direction)

-- VALUES
-- (
-- (SELECT id FROM applicants LIMIT 1),
-- (SELECT id FROM applications LIMIT 1),
-- 315.87,
-- '2025-08-15 00:00:00',
-- (SELECT id FROM payment_status WHERE identifier = 'completed'),
-- 'BaföG Geld',
-- 'DE89370400440532013000',
-- 'COBAFF',
-- 'incoming'
-- )

-- INSERT INTO payments
-- (applicant_id, application_id, amount, executed, status_id, 
-- 	description, iban, bic, direction)
-- VALUES
-- (
-- (SELECT id FROM applicants LIMIT 1),
-- (SELECT id FROM applications LIMIT 1),
-- 315.87,
-- null,
-- (SELECT id FROM payment_status WHERE identifier = 'pending'),
-- 'BaföG Geld',
-- 'DE89370400440532013000',
-- 'COBAFF',
-- 'incoming'
-- )