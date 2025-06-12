-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ready_products (
    id CHAR(26) PRIMARY KEY,
    kode CHAR(26) NOT NULL UNIQUE, 
    nama VARCHAR(255) NOT NULL,
    jumlah INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ready_products;
-- +goose StatementEnd