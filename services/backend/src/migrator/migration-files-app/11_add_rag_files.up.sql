CREATE TYPE rag_document_process_status_enum AS ENUM ();
ALTER TYPE rag_document_process_status_enum ADD VALUE 'IN_PROGRESS'; 
ALTER TYPE rag_document_process_status_enum ADD VALUE 'ERROR'; 
ALTER TYPE rag_document_process_status_enum ADD VALUE 'COMPLETED';

CREATE TABLE "rag_document_process_status" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "identifier" rag_document_process_status_enum NOT NULL UNIQUE,
  "name" varchar NOT NULL
);

INSERT INTO "rag_document_process_status" ("identifier", "name") VALUES 
('IN_PROGRESS', 'In Bearbeitung'),
('ERROR', 'Fehlerhaft'),
('COMPLETED', 'Abgeschlossen');


CREATE TABLE pgvector_rag_schuelerbafoeg_files (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  file_name VARCHAR(255) NOT NULL CHECK (file_name <> ''),
  file_type VARCHAR(50) NOT NULL CHECK (file_type <> ''),
  file_size BIGINT NOT NULL CHECK (file_size > 0),
  file_hash VARCHAR NOT NULL CHECK (file_hash <> ''),
  status UUID NOT NULL,
  processed_timestamp TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
  processed_error TEXT NOT NULL DEFAULT '',
  created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  "created_from" uuid NOT NULL
);

ALTER TABLE "pgvector_rag_schuelerbafoeg_files" ADD FOREIGN KEY ("status") REFERENCES "rag_document_process_status" ("id");
ALTER TABLE "pgvector_rag_schuelerbafoeg_files" ADD FOREIGN KEY ("created_from") REFERENCES "users" ("id");

CREATE TABLE pgvector_rag_studierendenbafoeg_files (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  file_name VARCHAR(255) NOT NULL CHECK (file_name <> ''),
  file_type VARCHAR(50) NOT NULL CHECK (file_type <> ''),
  file_size BIGINT NOT NULL CHECK (file_size > 0),
  file_hash VARCHAR NOT NULL CHECK (file_hash <> ''),
  status UUID NOT NULL,
  processed_timestamp TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
  processed_error TEXT NOT NULL DEFAULT '',
  created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  "created_from" uuid NOT NULL
);

CREATE TYPE rag_bafoeg_type_enum AS ENUM ();
ALTER TYPE rag_bafoeg_type_enum ADD VALUE 'SCHUELERBAFOEG';
ALTER TYPE rag_bafoeg_type_enum ADD VALUE 'STUDIERENDENBAFOEG';

CREATE TYPE rag_message_sender_enum AS ENUM ();
ALTER TYPE rag_message_sender_enum ADD VALUE 'USER';
ALTER TYPE rag_message_sender_enum ADD VALUE 'SYSTEM';

CREATE TABLE rag_conversations(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id),
  bafoeg_type rag_bafoeg_type_enum NOT NULL,
  created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE rag_conversation_messages(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  conversation_id UUID NOT NULL REFERENCES rag_conversations(id),
  message TEXT,
  sender rag_message_sender_enum NOT NULL,
  created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE VIEW rag_conversation_messages_overview AS (
  SELECT 
    rag_conversation_messages.id,
    rag_conversation_messages.conversation_id,
    rag_conversations.user_id,
    rag_conversations.bafoeg_type,
    rag_conversation_messages.message,
    rag_conversation_messages.sender,
    rag_conversation_messages.created
  FROM rag_conversation_messages
  INNER JOIN rag_conversations ON rag_conversation_messages.conversation_id = rag_conversations.id
);

ALTER TABLE "pgvector_rag_studierendenbafoeg_files" ADD FOREIGN KEY ("status") REFERENCES "rag_document_process_status" ("id");
ALTER TABLE "pgvector_rag_studierendenbafoeg_files" ADD FOREIGN KEY ("created_from") REFERENCES "users" ("id");

CREATE VIEW "pgvector_rag_schuelerbafoeg_files_overview" AS (
	SELECT 
		pgvector_rag_schuelerbafoeg_files.id,
		pgvector_rag_schuelerbafoeg_files.file_name,
		pgvector_rag_schuelerbafoeg_files.file_type,
		pgvector_rag_schuelerbafoeg_files.file_size,
		pgvector_rag_schuelerbafoeg_files.file_hash,

    json_build_object(
      'id', rag_document_process_status.id, 
      'identifier', rag_document_process_status.identifier, 
      'name', rag_document_process_status.name
    ) AS status,

    pgvector_rag_schuelerbafoeg_files.processed_timestamp,
    pgvector_rag_schuelerbafoeg_files.processed_error,

		pgvector_rag_schuelerbafoeg_files.created,
		
		json_build_object(
	        'id', users.id,
	        'username', users.username,
	        'display_name', users.display_name
	      ) AS "created_from"
	FROM pgvector_rag_schuelerbafoeg_files
	INNER JOIN users ON users.id = pgvector_rag_schuelerbafoeg_files.created_from
  INNER JOIN rag_document_process_status ON pgvector_rag_schuelerbafoeg_files.status = rag_document_process_status.id
);

CREATE VIEW "pgvector_rag_studierendenbafoeg_files_overview" AS (
	SELECT 
		pgvector_rag_studierendenbafoeg_files.id,
		pgvector_rag_studierendenbafoeg_files.file_name,
		pgvector_rag_studierendenbafoeg_files.file_type,
		pgvector_rag_studierendenbafoeg_files.file_size,
		pgvector_rag_studierendenbafoeg_files.file_hash,

    json_build_object(
      'id', rag_document_process_status.id, 
      'identifier', rag_document_process_status.identifier, 
      'name', rag_document_process_status.name
    ) AS status,

    pgvector_rag_studierendenbafoeg_files.processed_timestamp,
    pgvector_rag_studierendenbafoeg_files.processed_error,

		pgvector_rag_studierendenbafoeg_files.created,
		
		json_build_object(
	        'id', users.id,
	        'username', users.username,
	        'display_name', users.display_name
	      ) AS "created_from"
	FROM pgvector_rag_studierendenbafoeg_files
	INNER JOIN users ON users.id = pgvector_rag_studierendenbafoeg_files.created_from
  INNER JOIN rag_document_process_status ON pgvector_rag_studierendenbafoeg_files.status = rag_document_process_status.id
);
