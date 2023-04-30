-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE notification (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    section_id uuid NOT NULL,
    title varchar,
    description varchar,
    deadline time,
    status varchar,
    error_status int,
    PRIMARY KEY (id),
    FOREIGN KEY (section_id) REFERENCES section(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE notification;
