package repositories

import (
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/utils/types"
)

type GetFileReply struct {
	File models.File
}

type GetFileRequest struct {
	ID   string
	Type types.Types
}
