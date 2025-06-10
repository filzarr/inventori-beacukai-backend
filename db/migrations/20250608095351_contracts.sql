-- +goose Up
-- +goose StatementBegin

-- ENUM kategori_bc hanya dibuat jika belum ada (reuse dari bc_documents)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_barang') THEN
        CREATE TYPE kategori_barang AS ENUM ('Bahan Baku', 'Bahan Penolong', 'Mesin/Sparepart');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS contracts (
    id CHAR(26) PRIMARY KEY, 
    no_kontrak CHAR(26) NOT NULL,
    supliers_id CHAR(26) NOT NULL,
    kategori kategori_barang NOT NULL,
    tanggal TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT no_kontrak_unique UNIQUE (no_kontrak),
    FOREIGN KEY (supliers_id) REFERENCES supliers (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS contracts;
-- +goose StatementEnd