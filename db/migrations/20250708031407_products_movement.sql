-- +goose Up
-- +goose StatementBegin 
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_perpindahan') THEN
        CREATE TYPE status_perpindahan AS ENUM ('Diminta', 'Dikirim', 'Diterima');
    END IF;
END$$;
CREATE TABLE IF NOT EXISTS products_movement (
    id CHAR(26) PRIMARY KEY, 
    kode_barang CHAR(26) NOT NULL,
    jumlah INTEGER NOT NULL DEFAULT 0,
    status_perpindahan status_perpindahan NOT NULL DEFAULT 'Diminta',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (kode_barang) REFERENCES products (kode)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products_movement; 
-- +goose StatementEnd