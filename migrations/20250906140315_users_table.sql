-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT NOW(),
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(36) NOT NULL UNIQUE,
    password VARCHAR(250) NOT NULL,
    verified BOOLEAN DEFAULT false,
    blocked BOOLEAN DEFAULT false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
