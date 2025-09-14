package entity

import "inventori-beacukai-backend/pkg/types"

type GetPenyesuaianReq struct {
	Q          string `query:"q" validate:"omitempty"`
	KodeBarang string `query:"kode_barang" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetPenyesuaianReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Penyesuaian struct {
	Common
	KodeBarang string `json:"kode_barang" db:"kode_barang"`
	NamaBarang string `json:"nama_barang" db:"nama_barang"`
	Gudang     string `json:"gudang" db:"gudang"`
	Jumlah     int    `json:"jumlah" db:"jumlah"`
}

type GetPenyesuaianResp struct {
	Items []Penyesuaian `json:"items"`
	Meta  types.Meta    `json:"meta"`
}

type CreatePenyesuaianReq struct {
	KodeBarang string `json:"kode_barang" validate:"required,min=1"`
	Warehouse  string `json:"gudang" validate:"required,min=1"`
	Jumlah     int    `json:"jumlah" validate:"required,gt=0"`
}

type CreatePenyesuaianResp struct {
	Id string `json:"id"`
}
