-- +goose Up
-- +goose StatementBegin 
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_gudang') THEN
        CREATE TYPE kategori_gudang AS ENUM ('Bahan Baku', 'Bahan Penolong', 'Mesin/Sparepart', 'Barang Jadi', 'Produksi');
    END IF;
END$$;
CREATE TABLE IF NOT EXISTS warehouses (
    id CHAR(26) PRIMARY KEY,
    kode VARCHAR(26) NOT NULL UNIQUE, 
    nama VARCHAR(255) NOT NULL,
    kategori kategori_gudang NOT NULL,
    keterangan VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS warehouses; 
-- +goose StatementEnd