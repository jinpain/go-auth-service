-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    device TEXT,
    ip_address TEXT,
    created_at TIMESTAMP DEFAULT now(),
    revoked_at TIMESTAMP DEFAULT NULL,
    revoked BOOLEAN DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd
