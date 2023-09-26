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

type FileCompares struct {
	FileBeforeMerge []byte `json:"file_before_merge"`
	FileAfterMerge  []byte `json:"file_after_merge"`
}

type Changed struct {
	Added   []byte `json:"added"`
	Deleted []byte `json:"deleted"`
}

type FileMergeRequest struct {
	FileID      string       `json:"file_id"`
	FileCompare FileCompares `json:"file_compare"`
	FileChanged Changed      `json:"file_changed"`
}

type Files []FileMergeRequest

type MergeRequest struct {
	ID          uint64           `gorm:"primaryKey;autoIncrement:true"`
	Type        types.Types      `gorm:"type:integer;default:0;not null;index:idx_file_types"`
	File        Files            `gorm:"type:jsonb;default:'{}'"`
	Information MergeRequestInfo `gorm:"type:jsonb;default:'{}'"`
	CreatedAt   int64            `gorm:"autoCreateTime:nano;not null"`
	UpdatedAt   int64            `gorm:"autoUpdateTime:nano"`
	CreatedBy   string           `gorm:"type:text; not null"`
	UpdatedBy   string           `gorm:"type:text;"`
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

func (f *Files) Scan(src any) error {
	return ScanJSON(src, f)
}

func (f Files) Value() (driver.Value, error) {
	return json.Marshal(f)
}
