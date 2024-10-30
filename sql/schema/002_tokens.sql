-- +goose Up
CREATE TABLE tokens (
    id UUID PRIMARY KEY,
    email_id UUID REFERENCES emails(id) ON DELETE CASCADE,
    token TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    is_used BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose Down
DROP TABLE tokens;