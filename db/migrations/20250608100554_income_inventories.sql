-- +goose Up
-- +goose StatementBegin

-- ENUM kategori_barang_income
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_barang_income') THEN
        CREATE TYPE kategori_barang_income AS ENUM ('Bahan Baku', 'Bahan Penolong', 'Mesin/Sparepart');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS income_inventories (
    id CHAR(26) PRIMARY KEY,
    no_kontrak CHAR(26) NOT NULL,
    kategori_barang kategori_barang_income NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS income_inventories;
DROP TYPE IF EXISTS kategori_barang_income;

-- +goose StatementEnd