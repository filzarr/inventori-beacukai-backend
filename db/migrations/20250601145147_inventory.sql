-- +goose Up
-- +goose StatementBegin
/* ENUM ------------------------------------------------------------------- */
CREATE TYPE categories         AS ENUM (
  'Bahan Baku',
  'Bahan Penolong',
  'Barang Jadi',
  'Barang Sisa dan Scrap',
  'Mesin dan Peralatan Perkantoran'
);

CREATE TYPE inventories_status AS ENUM ('Dipinjam', 'Selesai', 'Tersedia');
CREATE TYPE document_status    AS ENUM ('Gate in TPB', 'Pembongkaran', 'Stuffing', 'Gate Out', 'Selesai');
CREATE TYPE document_type      AS ENUM ('BC 2.3', 'BC 2.5', 'BC 2.7 IN');

/* TABLE ------------------------------------------------------------------ */
CREATE TABLE IF NOT EXISTS inventories (
  id                CHAR(26)        PRIMARY KEY,
  inventories_status inventories_status NOT NULL DEFAULT 'Tersedia',
  kode_barang       CHAR(25)        NOT NULL,
  nama_barang       VARCHAR(255),
  kategori          categories      NOT NULL,
  pemasok           VARCHAR(255),
  pembeli           VARCHAR(255),
  saldo_awal        INTEGER         NOT NULL,
  satuan            CHAR(10)        NOT NULL,
  stok_opname       INTEGER         NOT NULL,
  mata_uang         CHAR(50)        NOT NULL,
  negara_asal       CHAR(50)        NOT NULL,
  document_type     document_type   NOT NULL,
  keterangan        VARCHAR(255),
  document_status   document_status NOT NULL, 
  created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at        TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS inventories;
DROP TYPE   IF EXISTS document_type;
DROP TYPE   IF EXISTS document_status;
DROP TYPE   IF EXISTS inventories_status;
DROP TYPE   IF EXISTS categories;
-- +goose StatementEnd