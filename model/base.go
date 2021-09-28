package model

import (
	"encoding/json"
)

type TimeHook struct {
	CreatedAt int64 `gorm:"not null" json:"-" swaggerignore:"true"`
	UpdatedAt int64 `gorm:"not null" json:"-" swaggerignore:"true"`
}
type AuthAccept struct {
	Error int             `json:"error"`
	Msg   string          `json:"msg"`
	Data  json.RawMessage `json:"data"`
}
