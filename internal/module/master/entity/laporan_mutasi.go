package entity

import "inventori-beacukai-backend/pkg/types"

type GetLaporanMutasiReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetLaporanMutasiReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type LaporanMutasi struct {
	KodeDokumen    string `json:"kode_dokumen" db:"kode_dokumen"`       // dari bc_documents.kategori
	NoDokumen      string `json:"no_dokumen" db:"no_dokumen"`           // dari bc_documents.no_document
	TanggalDokumen string `json:"tanggal_dokumen" db:"tanggal_dokumen"` // dari bc_documents.tanggal
	NoKontrak      string `json:"no_kontrak" db:"no_kontrak"`
	KodeBarang     string `json:"kode_barang" db:"kode_barang"`
	NamaBarang     string `json:"nama_barang" db:"nama_barang"`
	Satuan         string `json:"satuan" db:"satuan"`
	SaldoAwal      int    `json:"saldo_awal" db:"saldo_awal"`   // alias dari stok_awal di income_inventories_products
	Pemasukan      int    `json:"pemasukan" db:"pemasukan"`     // dari income_inventories_products.jumlah
	StokOpname     int    `json:"stok_opname" db:"stok_opname"` // dari transaction_incomes.jumlah
	StokAkhir      int    `json:"stok_akhir" db:"stok_akhir"`   // dari products.jumlah
}

type GetLaporanMutasiResp struct {
	Items []LaporanMutasi `json:"items"`
	Meta  types.Meta      `json:"meta"`
}
