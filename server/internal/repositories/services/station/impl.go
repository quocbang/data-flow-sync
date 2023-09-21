package station

import (
	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
)

type service struct {
	pg *gorm.DB
}

func NewService(pg *gorm.DB) repositories.StationServices {
	return service{
		pg: pg,
	}
}
