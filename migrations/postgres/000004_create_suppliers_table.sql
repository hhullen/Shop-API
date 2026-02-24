-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS suppliers (
    "uid" UUID NOT NULL,
    "name" TEXT NOT NULL,
    "phone_number" TEXT NOT NULL,
    address_id INT NOT NULL REFERENCES addresses(id) ON DELETE CASCADE,

    UNIQUE(uid)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS suppliers;

-- +goose StatementEnd
