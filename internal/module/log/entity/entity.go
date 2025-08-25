package entity

import (
	"database/sql"
	"encoding/json"
	"inventori-beacukai-backend/pkg/types"
	"time"
)

type GetLogReq struct {
	Q      string `query:"q" validate:"omitempty"`
	UserId string `json:"user_id"`
	types.MetaQuery
}

func (r *GetLogReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Log struct {
	Id        int64            `db:"id" json:"id"`
	TableName string           `db:"table_name" json:"table_name"`
	Operation string           `db:"operation" json:"operation"`
	UserID    *sql.NullString  `db:"user_id" json:"user_id"`
	OldData   *json.RawMessage `db:"old_data" json:"old_data,omitempty"`
	NewData   *json.RawMessage `db:"new_data" json:"new_data,omitempty"`
	ChangedAt time.Time        `db:"changed_at" json:"changed_at"`
}

type GetLogResp struct {
	Items []Log      `json:"items"`
	Meta  types.Meta `json:"meta"`
}
