-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS income_inventories_products (
    id CHAR(26) PRIMARY KEY,
    id_inventories CHAR(26) NOT NULL,
    kode_barang CHAR(26) NOT NULL,
    stok_awal INTEGER NOT NULL DEFAULT 0,
    jumlah INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,


    FOREIGN KEY (id_inventories) REFERENCES income_inventories (id),
    FOREIGN KEY (kode_barang) REFERENCES products (kode)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS income_inventories_products;
-- +goose StatementEnd
