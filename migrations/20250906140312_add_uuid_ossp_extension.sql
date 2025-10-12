-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'uuid-ossp') THEN
        CREATE EXTENSION "uuid-ossp" SCHEMA public;
    END IF;
END $$;
-- +goose StatementEnd