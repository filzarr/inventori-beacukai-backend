-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_produk') THEN
        CREATE TYPE kategori_produk AS ENUM ('Bahan Baku', 'Bahan Penolong', 'Mesin/Sparepart');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS products (
    id CHAR(26) PRIMARY KEY,
    kode CHAR(26) NOT NULL UNIQUE, 
    nama VARCHAR(255) NOT NULL,
    kategori kategori_produk NOT NULL,
    saldo_awal INTEGER NOT NULL DEFAULT 0,
    jumlah INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
DROP TYPE IF EXISTS kategori_produk;
-- +goose StatementEnd