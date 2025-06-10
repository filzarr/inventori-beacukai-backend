package entity

import "inventori-beacukai-backend/pkg/types"

type GetContractProductsReq struct {
	Q         string `query:"q" validate:"omitempty,min=3"`
	NoKontrak string `query:"noKontrak" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetContractProductsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type ContractProduct struct {
	Common
	NoKontrak      string `json:"no_kontrak" db:"no_kontrak"`
	KodeBarang     string `json:"kode_barang" db:"kode_barang"`
	NamaPemasok    string `json:"nama_pemasok" db:"nama_pemasok"`
	Alamat         string `json:"alamat" db:"alamat"`
	Stok           int    `json:"stok" db:"stok"`
	NamaBarang     string `json:"nama_barang" db:"nama_barang"`
	Jumlah         int    `json:"jumlah" db:"jumlah"`
	Satuan         string `json:"satuan" db:"satuan"`
	HargaSatuan    int    `json:"harga_satuan" db:"harga_satuan"`
	KodeMataUang   string `json:"kode_mata_uang" db:"kode_mata_uang"`
	NilaiBarangFog int    `json:"nilai_barang_fog" db:"nilai_barang_fog"`
	NilaiBarangRp  int    `json:"nilai_barang_rp" db:"nilai_barang_rp"`
}

type GetContractProductsResp struct {
	Items []ContractProduct `json:"items"`
	Meta  types.Meta        `json:"meta"`
}

type GetContractProductReq struct {
	Id string `json:"id" validate:"required"`
}

type GetContractProductResp struct {
	ContractProduct
}

type CreateContractProductReq struct {
	NoKontrak      string `json:"no_kontrak" validate:"required"`
	KodeBarang     string `json:"kode_barang" validate:"required"`
	Jumlah         int    `json:"jumlah" validate:"min=0"`
	Satuan         string `json:"satuan" validate:"required"`
	HargaSatuan    int    `json:"harga_satuan" validate:"min=0"`
	KodeMataUang   string `json:"kode_mata_uang" validate:"required"`
	NilaiBarangFog int    `json:"nilai_barang_fog" validate:"min=0"`
	NilaiBarangRp  int    `json:"nilai_barang_rp" validate:"min=0"`
}

type CreateContractProductResp struct {
	Id string `json:"id"`
}

type UpdateContractProductReq struct {
	Id             string `params:"id" validate:"required"`
	NoKontrak      string `json:"no_kontrak" validate:"required"`
	KodeBarang     string `json:"kode_barang" validate:"required"`
	Jumlah         int    `json:"jumlah" validate:"min=0"`
	Satuan         string `json:"satuan" validate:"required"`
	HargaSatuan    int    `json:"harga_satuan" validate:"min=0"`
	KodeMataUang   string `json:"kode_mata_uang" validate:"required"`
	NilaiBarangFog int    `json:"nilai_barang_fog" validate:"min=0"`
	NilaiBarangRp  int    `json:"nilai_barang_rp" validate:"min=0"`
}

type DeleteContractProductReq struct {
	Id string `json:"id" validate:"required"`
}
