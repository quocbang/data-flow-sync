package models

import "gorm.io/gorm/schema"

type Models interface {
	schema.Tabler
}

func GetModelList() []Models {
	return []Models{
		&Account{},
	}
}
