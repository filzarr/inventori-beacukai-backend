-- +goose Up
-- +goose StatementBegin

-- Buat enum type untuk kategori BC dan kategori barang jika belum ada
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_bc') THEN
        CREATE TYPE kategori_bc AS ENUM ('2.3', '2.5', '2.7', '2.61', '2.62', '40', '41', '30', 'PPFT2');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_document') THEN
        CREATE TYPE kategori_document AS ENUM ('Penjualan', 'Pembelian');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS bc_documents (
    id CHAR(26) PRIMARY KEY,
    kategori kategori_bc NOT NULL,
    kategori_document kategori_document NOT NULL DEFAULT 'Pembelian',
    no_document CHAR(50) NOT NULL,
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
DROP TYPE IF EXISTS kategori_document;
DROP TYPE IF EXISTS kategori_bc;

-- +goose StatementEnd