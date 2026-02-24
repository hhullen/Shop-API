-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS products (
    "uid" UUID NOT NULL,
    "name" TEXT NOT NULL,
    category TEXT NOT NULL,
    price BIGINT NOT NULL,
    available_stock BIGINT NOT NULL,
    last_update_date TIMESTAMPTZ NOT NULL,
    supplier_id UUID NOT NULL,
    image_id UUID NOT NULL,

    UNIQUE(uid)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS products;

-- +goose StatementEnd
