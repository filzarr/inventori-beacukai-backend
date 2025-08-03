-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id CHAR(26) PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id CHAR(26) NOT NULL,
    superadmin BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT email_unique UNIQUE (email),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
