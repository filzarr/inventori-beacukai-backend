-- +goose Up
-- +goose StatementBegin 
/* TABLE ------------------------------------------------------------------ */
CREATE TABLE
    IF NOT EXISTS pemasok_atau_pembeli (
        id CHAR(26) PRIMARY KEY,
        nama VARCHAR(255) NOT NULL,
        alamat VARCHAR(255) NOT NULL,
        noHp CHAR(255) NOT NULL,
        created_at TIMESTAMP
        WITH
            TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
            deleted_at TIMESTAMP
        WITH
            TIME ZONE
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS pemasok_atau_pembeli; 

-- +goose StatementEnd