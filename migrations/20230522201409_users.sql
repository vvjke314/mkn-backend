-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    username varchar UNIQUE,
    is_manager int,
    password varchar,
    email varchar UNIQUE,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE users;