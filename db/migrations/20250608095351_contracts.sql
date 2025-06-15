-- +goose Up
-- +goose StatementBegin

-- ENUM kategori_bc hanya dibuat jika belum ada (reuse dari bc_documents)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_kontrak') THEN
        CREATE TYPE kategori_kontrak AS ENUM ('Penjualan', 'Pembelian');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS contracts (
    id CHAR(26) PRIMARY KEY, 
    no_kontrak CHAR(26) NOT NULL,
    no_document CHAR(50), 
    kategori_kontrak kategori_kontrak NOT NULL DEFAULT 'Pembelian',
    supliers_id CHAR(26) NOT NULL,
    tanggal TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT no_kontrak_unique UNIQUE (no_kontrak),
    FOREIGN KEY (no_document) REFERENCES bc_documents (no_document),
    FOREIGN KEY (supliers_id) REFERENCES supliers (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS contracts;
DROP TYPE IF EXISTS kategori_kontrak;

-- +goose StatementEnd