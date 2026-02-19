-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS addresses (
    id SERIAL PRIMARY KEY,
    country TEXT NOT NULL,
    city TEXT NOT NULL,
    street TEXT NOT NULL,

    UNIQUE(country, city, street)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS addresses;

-- +goose StatementEnd