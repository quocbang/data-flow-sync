package models

import "github.com/lib/pq"

type Account struct {
	UserID   string        `gorm:"type:text,primaryKey"`
	Email    string        `gorm:"type:text;unique"`
	Password []byte        `gorm:"type:bytea; not null"`
	Roles    pq.Int64Array `gorm:"type:smallint[];default:'{}';not null"`
}

func (a *Account) TableName() string {
	return "account"
}
