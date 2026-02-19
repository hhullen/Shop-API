-- +goose Up
-- +goose StatementBegin

CREATE VIEW supplier_details AS
SELECT s.uid, s.name, s.phone_number, a.country, a.city, a.street
FROM suppliers s JOIN addresses a ON s.address_id = a.id;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP VIEW IF EXISTS supplier_details;

-- +goose StatementEnd