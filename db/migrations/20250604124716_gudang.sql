-- +goose Up
-- +goose StatementBegin 
CREATE TABLE IF NOT EXISTS gudang (
  id                CHAR(26)        PRIMARY KEY, 
  nama       VARCHAR(255)        NOT NULL,
  lokasi       VARCHAR(255)        NOT NULL,  
  user_id       CHAR(26) NOT NULL, 
  created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at        TIMESTAMP WITH TIME ZONE,
  FOREIGN KEY (user_id) REFERENCES users (id)

);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS gudang;  
-- +goose StatementEnd