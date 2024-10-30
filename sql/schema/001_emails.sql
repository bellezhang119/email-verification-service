-- +goose Up
CREATE TABLE emails (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email TEXT NOT NULL,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose Down
DROP TABLE emails;