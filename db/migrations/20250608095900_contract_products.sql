-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS contract_products (
    id CHAR(26) PRIMARY KEY,
    no_kontrak CHAR(26) NOT NULL,
    kode_barang CHAR(26) NOT NULL, -- disamakan dengan products.kode
    jumlah INTEGER NOT NULL DEFAULT 0,
    satuan CHAR(50) NOT NULL,
    harga_satuan BIGINT NOT NULL DEFAULT 0,
    kode_mata_uang CHAR(50) NOT NULL,
    nilai_barang_fog BIGINT NOT NULL DEFAULT 0,
    nilai_barang_rp BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (no_kontrak) REFERENCES contracts (no_kontrak),
    FOREIGN KEY (kode_barang) REFERENCES products (kode) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS contract_products;
-- +goose StatementEnd