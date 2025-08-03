-- +goose Up
-- +goose StatementBegin

-- Buat enum type untuk kategori BC dan kategori barang jika belum ada 

CREATE TABLE IF NOT EXISTS bc_documents (
    id CHAR(26) PRIMARY KEY,
    kategori VARCHAR(50) NOT NULL,
    kode_document VARCHAR(50) NOT NULL, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT kode_document UNIQUE (kode_document)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS bc_documents;

-- +goose StatementEnd