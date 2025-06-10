-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS currencies (
    id char(26) PRIMARY KEY,
    kode CHAR(26) NOT NULL,
    mata_uang char(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,


    CONSTRAINT kode_unique UNIQUE (kode)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS currencies;
-- +goose StatementEnd
