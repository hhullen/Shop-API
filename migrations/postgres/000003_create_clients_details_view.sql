-- +goose Up
-- +goose StatementBegin

CREATE VIEW client_details AS
SELECT 
    c.client_name, c.client_surname, c.birthday, c.gender,
    c.uid, c.registration_date, a.country, a.city, a.street
FROM clients c JOIN addresses a ON c.address_id = a.id;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP VIEW IF EXISTS client_details;

-- +goose StatementEnd