-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS income_inventories_products (
    id CHAR(26) PRIMARY KEY,
    no_kontrak CHAR(26) NOT NULL,
    kode_barang CHAR(26) NOT NULL,
    stok_awal INTEGER NOT NULL DEFAULT 0,
    lokasi CHAR(50) NOT NULL,
    jumlah INTEGER NOT NULL DEFAULT 0, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,


    FOREIGN KEY (no_kontrak) REFERENCES contracts (no_kontrak),
    FOREIGN KEY (kode_barang) REFERENCES products (kode)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS income_inventories_products;
-- +goose StatementEnd
