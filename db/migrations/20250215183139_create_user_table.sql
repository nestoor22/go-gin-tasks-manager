-- +goose Up
-- +goose StatementBegin
DROP
EXTENSION IF EXISTS "uuid-ossp";
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id         UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    first_name VARCHAR                                                 NOT NULL,
    last_name  VARCHAR                                                 NOT NULL,
    email      VARCHAR UNIQUE                                          NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc', now()) NOT NULL,
    update_at  TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc', now()) NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
