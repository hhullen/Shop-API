-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS images (
    "uid" UUID NOT NULL,
    "image" BYTEA NOT NULL,

    UNIQUE(uid)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS images;

-- +goose StatementEnd
