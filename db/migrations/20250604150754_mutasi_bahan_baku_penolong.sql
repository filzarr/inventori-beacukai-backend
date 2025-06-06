-- +goose Up
-- +goose StatementBegin
CREATE TYPE status AS ENUM ('Diajukan', 'Proses', 'Diterima');

/* ENUM ------------------------------------------------------------------- */
CREATE TABLE
    IF NOT EXISTS mutasi_bahan_baku_penolong (
        id CHAR(26) PRIMARY KEY,
        inventories_id CHAR(26) NOT NULL,
        gudang_id CHAR(26) NOT NULL,
        jumlah INTEGER,
        saldo_awal INTEGER DEFAULT 0,
        pemasukan INTEGER DEFAULT 0,
        pengeluaran INTEGER DEFAULT 0,
        penyesuaian INTEGER DEFAULT 0,
        status status NOT NULL DEFAULT 'Diajukan',
        created_at TIMESTAMP
        WITH
            TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
            deleted_at TIMESTAMP
        WITH
            TIME ZONE ,
        FOREIGN KEY (inventories_id) REFERENCES inventories (id),
        FOREIGN KEY (gudang_id) REFERENCES gudang (id)

    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS mutasi_bahan_baku_penolong;

DROP TYPE IF EXISTS status;

-- +goose StatementEnd