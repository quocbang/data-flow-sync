package models

type File struct {
	ID           string `gorm:"type:text;primaryKey"`
	FileContent  []byte `gorm:"type:bytea"`
	CreatedAt    int64  `gorm:"autoCreateTime:nano;not null"`
	LastMergedAt int64  `gorm:"autoUpdateTime:nano"`
	CreatedBy    string `gorm:"type:text"`
}

func (File) TableName() string {
	return "file"
}
