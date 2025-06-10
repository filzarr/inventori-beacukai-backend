-- +goose Up
-- +goose StatementBegin

-- ENUM kategori_barang_income

CREATE TABLE IF NOT EXISTS transaction_incomes (
    id CHAR(26) PRIMARY KEY,
    no_document CHAR(50) NOT NULL,
    no_kontrak CHAR(26) NOT NULL,
    kode_barang CHAR(26) NOT NULL,
    jumlah int NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (no_document) REFERENCES bc_documents (no_document),
    FOREIGN KEY (no_kontrak) REFERENCES contracts (no_kontrak),
    FOREIGN KEY (kode_barang) REFERENCES products (kode)

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS transaction_incomes;

-- +goose StatementEnd