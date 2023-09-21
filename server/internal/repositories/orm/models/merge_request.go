package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/lib/pq"

	"github.com/quocbang/data-flow-sync/server/utils/types"
)

type MergeRequestStatus struct {
	IsApproved bool `json:"is_approved"`
	IsMerged   bool `json:"is_merged"`
	IsOpen     bool `json:"is_open"`
}

type MergeRequestInfo struct {
	HistoryChanged pq.StringArray     `json:"history_changed"`
	Status         MergeRequestStatus `json:"status"`
}

type MergeRequest struct {
	ID              uint64           `gorm:"primaryKey;autoIncrement:true"`
	FileID          string           `gorm:"type:text;not null"`
	Type            types.Types      `gorm:"type:integer;default:0;not null;index:idx_file_types"`
	FileBeforeMerge []byte           `gorm:"type:bytea"`
	FileAfterMerge  []byte           `gorm:"type:bytea"`
	Information     MergeRequestInfo `gorm:"type:jsonb;default:'{}'"`
	CreatedAt       int64            `gorm:"autoCreateTime:nano;not null"`
	UpdatedAt       int64            `gorm:"autoUpdateTime:nano"`
	CreatedBy       string           `gorm:"type:text; not null"`
	UpdatedBy       string           `gorm:"type:text; not null"`
}

func (MergeRequest) TableName() string {
	return "merge_request"
}

func (m *MergeRequestInfo) Scan(src any) error {
	return ScanJSON(src, m)
}

func (m MergeRequest) Value() (driver.Value, error) {
	return json.Marshal(m)
}
