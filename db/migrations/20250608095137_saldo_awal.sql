-- +goose Up
-- +goose StatementBegin

-- Buat enum type untuk kategori BC jika belum ada
CREATE TABLE IF NOT EXISTS saldo_awals (
    id CHAR(26) PRIMARY KEY,
    kode_barang VARCHAR(26) NOT NULL,
    saldo_awal VARCHAR(26) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (kode_barang) REFERENCES products (kode)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS saldo_awals 
-- +goose StatementEnd