-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS contracts_bc (
    id CHAR(26) PRIMARY KEY,
    no_kontrak VARCHAR(26) NOT NULL,
    kode_document_bc VARCHAR(26) NOT NULL,
    nomor_document_bc VARCHAR(26) NOT NULL,
    tanggal_document_bc TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,  
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    FOREIGN KEY (no_kontrak) REFERENCES contracts (no_kontrak),
    FOREIGN KEY (kode_document_bc) REFERENCES bc_documents (kode_document) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS contracts_bc;
-- +goose StatementEnd