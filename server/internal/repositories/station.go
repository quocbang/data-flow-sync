package repositories

import "github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"

// CreateMRRequest definition.
type CreateMRRequest struct {
	MergeRequest models.MergeRequest
}

// CreateMRReply definition.
type CreateMRReply struct {
	MergeRequestID int64 `json:"merge_request_id"`
}

// GetMRRequest definition.
type GetMRRequest struct {
	MergeRequestID int64 `json:"merge_request_id"`
}

// GetMRReply definition.
type GetMRReply struct {
	MergeRequest models.MergeRequest
}

// Station definition.
type Station struct {
	ID           string  `json:"ID,omitempty" yaml:"ID,omitempty"`
	SubCompany   int64   `json:"sub_company" yaml:"sub-company,omitempty"`
	Factory      string  `json:"factory" yaml:"factory,omitempty"`
	DepartmentID string  `json:"department_id" yaml:"department-id,omitempty"`
	Alias        string  `json:"alias" yaml:"alias,omitempty"`
	SerialNumber int64   `json:"serial_number" yaml:"serial-number,omitempty"`
	Description  string  `json:"description" yaml:"description,omitempty"`
	Devices      []int64 `json:"devices" yaml:"devices,omitempty"`
}
