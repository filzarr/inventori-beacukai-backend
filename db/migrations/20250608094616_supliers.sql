-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_supliers') THEN
        CREATE TYPE kategori_supliers AS ENUM ('Penjual', 'Pembeli');
    END IF;
END$$;
CREATE TABLE IF NOT EXISTS supliers (
    id char(26) PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    kategori_supliers kategori_supliers NOT NULL,
    npwp VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS supliers;
-- +goose StatementEnd
