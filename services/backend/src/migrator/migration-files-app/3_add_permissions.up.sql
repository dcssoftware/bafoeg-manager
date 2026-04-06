CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    identifier VARCHAR(255) UNIQUE NOT NULL,
    inherit_from UUID NULL,
    FOREIGN KEY (inherit_from) REFERENCES roles(id) ON DELETE SET NULL,
    CHECK (inherit_from != id)
);

CREATE TABLE IF NOT EXISTS user_permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(255) NOT NULL,
    identifier VARCHAR(255) UNIQUE NOT NULL
    CHECK (identifier != '')
);

CREATE TABLE IF NOT EXISTS role_permissions (
    role_id UUID,
    permission_id UUID,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES user_permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_id UUID,
    role_id UUID,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

CREATE OR REPLACE FUNCTION get_user_permissions(p_user_id UUID) 
RETURNS JSONB AS $$
BEGIN
    RETURN (
        WITH RECURSIVE role_hierarchy AS (
            -- Base case: Get all roles directly assigned to the user
            SELECT
                ur.user_id,
                r.id AS role_id
            FROM    user_roles ur
            JOIN    roles r ON ur.role_id = r.id
            WHERE   ur.user_id = p_user_id
            UNION
            
            -- Recursive case: Get all roles that are inherited from other roles
            SELECT
                rh.user_id,
                r2.id AS role_id
            FROM    role_hierarchy rh
            JOIN    roles r1 ON rh.role_id = r1.id
            JOIN    roles r2 ON r1.inherit_from = r2.id
        )
        SELECT 
            jsonb_agg(DISTINCT p.identifier) AS permissions
        FROM    role_hierarchy rh
        JOIN    role_permissions rp ON rh.role_id = rp.role_id
        JOIN    user_permissions p ON rp.permission_id = p.id
    );
END;
$$ LANGUAGE plpgsql;

INSERT INTO "user_permissions" (name, description, identifier) VALUES 
( 'User Management - Read Users', 'Can view the user management tab and see all registered users', 'read:user-management'),
( 'RAG Management Schüler - Read Files', 'Can view the RAG Management Schüler files', 'read:rag-management-schueler-files'),
( 'RAG Management Schüler - Upload Files', 'Can upload files to the RAG Management Schüler', 'upload:rag-management-schueler-files'),
( 'RAG Management Schüler - Delete Files', 'Can delete files from the RAG Management Schüler', 'delete:rag-management-schueler-files'),
( 'RAG Management Studierenden - Read Files', 'Can view the RAG Management Studierenden files', 'read:rag-management-studierenden-files'),
( 'RAG Management Studierenden - Upload Files', 'Can upload files to the RAG Management Studierenden', 'upload:rag-management-studierenden-files'),
( 'RAG Management Studierenden - Delete Files', 'Can delete files from the RAG Management Studierenden', 'delete:rag-management-studierenden-files');

INSERT INTO "roles" (name, identifier, inherit_from) VALUES (
    'Admin', 'admin', NULL
);

INSERT INTO "role_permissions" (role_id, permission_id) VALUES 
(
        (SELECT id FROM roles WHERE name = 'Admin'),
        (SELECT id FROM user_permissions WHERE identifier = 'read:user-management')
),
(
        (SELECT id FROM roles WHERE name = 'Admin'),
        (SELECT id FROM user_permissions WHERE identifier = 'read:rag-management-schueler-files')
),
(
        (SELECT id FROM roles WHERE name = 'Admin'),
        (SELECT id FROM user_permissions WHERE identifier = 'upload:rag-management-schueler-files')
),
(
        (SELECT id FROM roles WHERE name = 'Admin'),
        (SELECT id FROM user_permissions WHERE identifier = 'delete:rag-management-schueler-files')
),
(
        (SELECT id FROM roles WHERE name = 'Admin'),
        (SELECT id FROM user_permissions WHERE identifier = 'read:rag-management-studierenden-files')
),
(
        (SELECT id FROM roles WHERE name = 'Admin'),
        (SELECT id FROM user_permissions WHERE identifier = 'upload:rag-management-studierenden-files')
),
(
        (SELECT id FROM roles WHERE name = 'Admin'),
        (SELECT id FROM user_permissions WHERE identifier = 'delete:rag-management-studierenden-files')
);
