-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE notifications (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    section_id uuid NOT NULL,
    title varchar,
    description varchar,
    deadline timestamp,
    status varchar,
    error_status int,
    PRIMARY KEY (id),
    FOREIGN KEY (section_id) REFERENCES sections(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE notifications;
