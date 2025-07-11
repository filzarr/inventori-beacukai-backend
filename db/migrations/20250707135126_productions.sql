-- +goose Up
-- +goose StatementBegin 

CREATE TABLE IF NOT EXISTS productions (
    id CHAR(26) PRIMARY KEY, 
    kode_barang CHAR(26) NOT NULL,
    jumlah INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (kode_barang) REFERENCES products (kode)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS productions; 
-- +goose StatementEnd