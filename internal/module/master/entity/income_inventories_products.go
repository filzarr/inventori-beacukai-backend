package entity

import "inventori-beacukai-backend/pkg/types"

type GetIncomeInventoryProductsReq struct {
	Q        string `query:"q" validate:"omitempty,min=3"`
	Full     bool   `query:"full" validate:"omitempty"`
	Kategori string `query:"kategori" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetIncomeInventoryProductsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type IncomeInventoryProduct struct {
	Common
	NoKontrak       string `json:"no_kontrak" db:"no_kontrak"`
	JumlahRealisasi int    `json:"jumlah_realisasi" db:"jumlah_realisasi"`
	JumlahKontrak   int    `json:"jumlah_kontrak" db:"jumlah_kontrak"`
	KodeBarang      string `json:"kode_barang" db:"kode_barang"`
	NamaBarang      string `json:"nama_barang" db:"nama_barang"`
	NilaiBarangFog  int    `json:"nilai_barang_fog" db:"nilai_barang_fog"`
	NilaiBarangRp   int    `json:"nilai_barang_rp" db:"nilai_barang_rp"`
	SaldoAwal       int    `json:"saldo_awal" db:"saldo_awal"`
}

type GetIncomeInventoryProductsResp struct {
	Items []IncomeInventoryProduct `json:"items"`
	Meta  types.Meta               `json:"meta"`
}

type GetIncomeInventoryProductReq struct {
	Id string `json:"id" validate:"required"`
}

type GetIncomeInventoryProductResp struct {
	IncomeInventoryProduct
}

type CreateIncomeInventoryProductReq struct {
	NoKontrak         string  `json:"no_kontrak" validate:"required"`
	KodeBarang        string  `json:"kode_barang" validate:"required"`
	WarehouseLocation *string `json:"warehouse_location"`
	Driver            string  `json:"driver" validate:"required"`
	LicensePlate      string  `json:"license_plate" validate:"required"`
	BrutoWeight       int64   `json:"bruto_weight" validate:"required"`
	NettoWeight       int64   `json:"netto_weight" validate:"required"`
	EmptyWeight       int64   `json:"empty_weight" validate:"required"`
	StartingTime      string  `json:"starting_time" validate:"required"`
	EndingTime        string  `json:"ending_time" validate:"required"`
	Tanggal           string  `json:"tanggal" validate:"required"`
	SaldoAwal         int     `json:"saldo_awal" validate:"required"`
	Jumlah            int     `json:"jumlah" validate:"min=0"`
}

type CreateIncomeInventoryProductResp struct {
	Id string `json:"id"`
}

type UpdateIncomeInventoryProductReq struct {
	Id                string `params:"id" validate:"required"`
	NoKontrak         string `json:"no_kontrak" validate:"required"`
	KodeBarang        string `json:"kode_barang" validate:"required"`
	WarehouseLocation string `json:"warehouse_location" validate:"required"`
	Driver            string `json:"driver" validate:"required"`
	LicensePlate      string `json:"license_plate" validate:"required"`
	BrutoWeight       int64  `json:"bruto_weight" validate:"required"`
	EmptyWeight       int64  `json:"empty_weight" validate:"required"`
	NettoWeight       int64  `json:"netto_weight" validate:"required"`
	StartingTime      string `json:"starting_time" validate:"required"`
	EndingTime        string `json:"ending_time" validate:"required"`
	SaldoAwal         string `json:"saldo_awal" db:"saldo_awal"`
	Jumlah            int    `json:"jumlah" validate:"min=0"`
}

type DeleteIncomeInventoryProductReq struct {
	Id string `json:"id" validate:"required"`
}

type GetIncomeInventoryProductsByContractReq struct {
	NoKontrak string `query:"no_kontrak" validate:"required"`
	types.MetaQuery
}

type IncomeInventoryProductByContract struct {
	Id              string `json:"id" db:"id"`
	NoKontrak       string `json:"no_kontrak" db:"no_kontrak"`
	JumlahRealisasi int    `json:"jumlah_realisasi" db:"jumlah_realisasi"`
	JumlahKontrak   int    `json:"jumlah_kontrak" db:"jumlah_kontrak"`
	KodeBarang      string `json:"kode_barang" db:"kode_barang"`
	NamaBarang      string `json:"nama_barang" db:"nama_barang"`
	NilaiBarangFog  int    `json:"nilai_barang_fog" db:"nilai_barang_fog"`
	NilaiBarangRp   int    `json:"nilai_barang_rp" db:"nilai_barang_rp"`
	SaldoAwal       int    `json:"saldo_awal" db:"saldo_awal"`
}
type GetIncomeInventoryProductsByContractResp struct {
	Items []IncomeInventoryProductByContract `json:"items"`
	Meta  types.Meta                         `json:"meta"`
}

func (r *GetIncomeInventoryProductsByContractReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type GetIncomeInventoryProductsByContractAndKodeReq struct {
	NoKontrak  string `query:"no_kontrak" validate:"required"`
	KodeBarang string `query:"kode_barang" validate:"required"`
	types.MetaQuery
}

type IncomeInventoryProductsByContractAndKode struct {
	Id                string `json:"id" db:"id"`
	NoKontrak         string `json:"no_kontrak" db:"no_kontrak"`
	JumlahMasuk       int    `json:"jumlah_masuk" db:"jumlah_masuk"`
	KodeBarang        string `json:"kode_barang" db:"kode_barang"`
	NamaBarang        string `json:"nama_barang" db:"nama_barang"`
	LokasiPenyimpanan string `json:"lokasi_penyimpanan" db:"lokasi_penyimpanan"`
	Tanggal           string `json:"tanggal" db:"tanggal"`
	JamMasuk          string `json:"jam_masuk" db:"jam_masuk"`
	JamKeluar         string `json:"jam_keluar" db:"jam_keluar"`
	Driver            string `json:"driver" db:"driver"`
	LicensePlate      string `json:"license_plate" db:"license_plate"`
	BrutoWeight       int64  `json:"bruto_weight" db:"bruto_weight"`
	NettoWeight       int64  `json:"netto_weight" db:"netto_weight"`
}

type GetIncomeInventoryProductsByContractAndKodeResp struct {
	Items []IncomeInventoryProductsByContractAndKode `json:"items"`
	Meta  types.Meta                                 `json:"meta"`
}

func (r *GetIncomeInventoryProductsByContractAndKodeReq) SetDefault() {
	r.MetaQuery.SetDefault()
}
