-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE collaborations (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    project_id uuid NOT NULL,
    user_id uuid NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE collaborations;
