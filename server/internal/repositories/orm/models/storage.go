package models

import (
	type_ "github.com/quocbang/data-flow-sync/server/assets/protobuf/types"
)

type DataStorage struct {
	ID        string     `gorm:"type:text,primaryKey"`
	Type      type_.Type `gorm:"not null"`
	Content   []byte     `gorm:"type:bytea; not null"`
	CreatedBy string     `gorm:"type:text, not null"`
	CreatedAt int64      `gorm:"type:bigint;autoCreateTime"`
}

func (DataStorage) Table() string {
	return "data_storage"
}
