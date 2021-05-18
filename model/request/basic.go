package request

type Basic struct {
	Timestamp int `json:"timestamp" binding:"required" validate:"required"`
	Token string `json:"token" binding:"required" validate:"required"`
	Data map[string]interface{} `json:"data" binding:"required" validate:"required"`
}