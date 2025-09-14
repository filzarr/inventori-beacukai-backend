-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS penyesuaian (
    id CHAR(26) PRIMARY KEY, 
    kode_barang VARCHAR(26) NOT NULL,
    warehouse_kode VARCHAR(26) NOT NULL,
    jumlah INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
  
    FOREIGN KEY (kode_barang) REFERENCES products (kode),
    FOREIGN KEY (warehouse_kode) REFERENCES warehouses (kode)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
