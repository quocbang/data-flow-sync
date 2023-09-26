package models

import "github.com/quocbang/data-flow-sync/server/utils/types"

type File struct {
	ID           string      `gorm:"type:text;primaryKey"`
	Type         types.Types `gorm:"type:integer;index:idx_file_type"`
	FileContent  []byte      `gorm:"type:bytea"`
	CreatedAt    int64       `gorm:"autoCreateTime:nano;not null"`
	LastMergedAt int64       `gorm:"autoUpdateTime:nano"`
	CreatedBy    string      `gorm:"type:text"`
}

func (File) TableName() string {
	return "file"
}
