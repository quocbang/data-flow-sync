package repositories

type CreateStationMRReply struct {
	MergeRequestID int64 `json:"merge_request_id"`
}

type GetMRRequest struct {
	MergeRequestID int64 `json:"merge_request_id"`
}

type GetMRReply struct {
	// TODO: get in models
}
