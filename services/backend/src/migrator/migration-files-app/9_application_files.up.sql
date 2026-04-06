CREATE TYPE file_content_type AS ENUM ('FORMBLATT01', 'FORMBLATT02', 'FORMBLATT03', 'FORMBLATT04', 'FORMBLATT05', 'FORMBLATT06', 'FORMBLATT07', 'FORMBLATT08', 'FORMBLATT09', 'FORMBLATT10', 'OTHER');

CREATE TABLE files (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  file_name VARCHAR(255) NOT NULL CHECK (file_name <> ''),
  file_type VARCHAR(50) NOT NULL CHECK (file_type <> ''),
  file_size BIGINT NOT NULL CHECK (file_size > 0),
  file_hash VARCHAR NOT NULL CHECK (file_hash <> ''),
  file_content_type file_content_type NOT NULL DEFAULT 'OTHER',
  file_content_ocr TEXT NOT NULL DEFAULT '',
  created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE application_files (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  application_id UUID NOT NULL,
  file_id UUID NOT NULL REFERENCES files(id),
  created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

ALTER TABLE application_files ADD FOREIGN KEY (application_id) REFERENCES applications (id);

CREATE INDEX idx_application_files_application_id ON application_files (application_id);
