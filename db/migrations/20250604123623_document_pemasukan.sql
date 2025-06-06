-- +goose Up
-- +goose StatementBegin
/* ENUM ------------------------------------------------------------------- */
CREATE TYPE document_type AS ENUM ('BC 2.3', 'BC 2.5', 'BC 2.7 IN');

/* TABLE ------------------------------------------------------------------ */
CREATE TABLE
    IF NOT EXISTS document_pemasukan (
        id CHAR(26) PRIMARY KEY,
        no_kontrak VARCHAR(255),
        document_type document_type NOT NULL,
        inventories_id CHAR(26) NOT NULL,
        keterangan VARCHAR(255),
        pemasok_id CHAR(26) NOT NULL,
        document TEXT NOT NULL,
        saldo_awal INTEGER NOT NULL, 
        nama_penandatangan CHAR(50) NOT NULL,
        tanggal TIMESTAMP
        WITH
            TIME ZONE NOT NULL,
            created_at TIMESTAMP
        WITH
            TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
            deleted_at TIMESTAMP
        WITH
            TIME ZONE,
            FOREIGN KEY (inventories_id) REFERENCES inventories (id),
            FOREIGN KEY (pemasok_id) REFERENCES pemasok_atau_pembeli (id),
            CONSTRAINT no_kontrak_unique UNIQUE (no_kontrak)
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin 
DROP TABLE IF EXISTS document_pemasukan;
-- DROP TYPE   IF EXISTS document_type;

-- +goose StatementEnd