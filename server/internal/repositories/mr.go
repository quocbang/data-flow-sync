package repositories

import "github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"

type GetMergeRequestOpeningByFileIDReply struct {
	MergeRequest models.MergeRequest
}
