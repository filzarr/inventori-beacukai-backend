-- +goose Up
-- +goose StatementBegin

-- Buat enum type untuk kategori BC dan kategori barang jika belum ada
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_bc') THEN
        CREATE TYPE kategori_bc AS ENUM ('BC 23', 'BC 27 In', 'BC 262', 'BC 40');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_barang') THEN
        CREATE TYPE kategori_barang AS ENUM ('Bahan Baku', 'Bahan Penolong', 'Mesin/Sparepart');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS bc_documents (
    id CHAR(26) PRIMARY KEY,
    kategori kategori_bc NOT NULL,
    no_document CHAR(50) NOT NULL,
    kategori_barang kategori_barang NOT NULL,
    tanggal TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT no_document_unique UNIQUE (no_document)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS bc_documents;
DROP TYPE IF EXISTS kategori_barang;
DROP TYPE IF EXISTS kategori_bc;

-- +goose StatementEnd