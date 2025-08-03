-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transfers_products (
    id char(26) PRIMARY KEY,
    kode_barang VARCHAR(26) NOT NULL, 
    jumlah VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (kode_barang) REFERENCES products (kode)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transfers_products;
-- +goose StatementEnd
