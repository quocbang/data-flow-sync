package models

import "github.com/lib/pq"

type Account struct {
	UserID   string        `gorm:"type:text,primaryKey"`
	Password []byte        `gorm:"type:bytea; not null"`
	Roles    pq.Int64Array `gorm:"type:smallint[];default:'{}';not null"`
}

func (a *Account) TableName() string {
	return "account"
}
