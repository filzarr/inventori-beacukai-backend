-- +goose Up
-- +goose StatementBegin

-- ENUM kategori_bc hanya dibuat jika belum ada (reuse dari bc_documents) 

CREATE TABLE IF NOT EXISTS warehouses_stocks (
    id CHAR(26) PRIMARY KEY, 
    warehouse_kode VARCHAR(26) NOT NULL,
    kode_barang VARCHAR(26) NOT NULL,
    jumlah INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (warehouse_kode) REFERENCES warehouses (kode),
    FOREIGN KEY (kode_barang) REFERENCES products (kode)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS warehouses_stocks;

-- +goose StatementEnd