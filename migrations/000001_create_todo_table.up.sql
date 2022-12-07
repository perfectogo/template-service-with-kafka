CREATE TYPE PRIORITY_TYPE AS ENUM ('high', 'meduim', 'low');

CREATE TABLE todos (
    id uuid NOT NULL PRIMARY KEY,
    name CHARACTER VARYING NOT NULL,
    priority PRIORITY_TYPE DEFAULT 'low',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)

CREATE UNIQUE INDEX todo_id_index ON todos (id);
CREATE INDEX ON todos((id));
