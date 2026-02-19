-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS clients (
    "uid" UUID NOT NULL,
    client_name TEXT NOT NULL,
    client_surname TEXT NOT NULL,
    birthday TIMESTAMPTZ NOT NULL,
    gender TEXT NOT NULL,
    registration_date TIMESTAMPTZ NOT NULL,
    address_id INT NOT NULL REFERENCES addresses(id) ON DELETE CASCADE,

    UNIQUE(client_name, client_surname, birthday, gender, registration_date, address_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS clients;

-- +goose StatementEnd