package models

type Station struct {
	ID           string  `json:"ID"`
	SubCompany   int64   `json:"sub_company"`
	Factory      string  `json:"factory"`
	DepartmentID string  `json:"department_id"`
	Alias        string  `json:"alias"`
	SerialNumber int64   `json:"serial_number"`
	Description  string  `json:"description"`
	Devices      []int64 `json:"devices"`
}
