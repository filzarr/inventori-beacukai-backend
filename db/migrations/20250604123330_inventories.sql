-- +goose Up
-- +goose StatementBegin
/* ENUM ------------------------------------------------------------------- */
CREATE TYPE categories         AS ENUM (
  'Bahan Baku',
  'Bahan Penolong',
  'Sparepart',
  'Sisa Hasil Produksi/Scrap',
  'Mesin dan Peralatan'
);

CREATE TYPE inventories_status AS ENUM ('Dipinjam', 'Selesai', 'Tersedia'); 

/* TABLE ------------------------------------------------------------------ */
CREATE TABLE IF NOT EXISTS inventories (
  id                CHAR(26)        PRIMARY KEY,
  inventories_status inventories_status NOT NULL DEFAULT 'Tersedia',
  kode_barang       CHAR(25)        NOT NULL,
  nama_barang       VARCHAR(255),
  kategori          categories      NOT NULL,  
  jumlah            INTEGER         NOT NULL DEFAULT 0,
  created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at        TIMESTAMP WITH TIME ZONE,
  CONSTRAINT kode_barang_unique UNIQUE (kode_barang)

);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS inventories; 
DROP TYPE   IF EXISTS inventories_status;
DROP TYPE   IF EXISTS categories;
-- +goose StatementEnd