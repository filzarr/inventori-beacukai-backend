-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS buyers (
    id char(26) PRIMARY KEY,
    name CHAR(25) NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    npwn char(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS buyers;
-- +goose StatementEnd
