-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS income_inventories_products (
    id CHAR(26) PRIMARY KEY,
    no_kontrak VARCHAR(26) NOT NULL,
    kode_barang VARCHAR(26) NOT NULL,
    stok_awal INTEGER NOT NULL DEFAULT 0,
    warehouse_location VARCHAR(26) DEFAULT NULL,
    driver VARCHAR(50) NOT NULL,
    license_plate VARCHAR(50) NOT NULL,
    bruto_weight BIGINT NOT NULL DEFAULT 0,
    empty_weight BIGINT NOT NULL DEFAULT 0,
    netto_weight BIGINT NOT NULL DEFAULT 0,
    starting_time TIME NOT NULL,
    ending_time TIME NOT NULL,
    jumlah INTEGER NOT NULL DEFAULT 0, 
    tanggal TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (warehouse_location) REFERENCES warehouses (kode),
    FOREIGN KEY (no_kontrak) REFERENCES contracts (no_kontrak),
    FOREIGN KEY (kode_barang) REFERENCES products (kode)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS income_inventories_products;
-- +goose StatementEnd
