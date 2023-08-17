package models

type Station struct {
	ID        string `gorm:"type:text;primaryKey"`
	Content   []byte `gorm:"type:bytea;not null"`
	CreatedBy string `gorm:"type:text;not null"`
	CreatedAt int64  `gorm:"type:bigint;autoCreateTime"`
}
