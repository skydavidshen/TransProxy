package request

type Basic struct {
	Timestamp uint64 `json:"timestamp" binding:"required"`
	Token string `json:"token" binding:"required"`
	Data map[string]interface{} `json:"data" binding:"required"`
}