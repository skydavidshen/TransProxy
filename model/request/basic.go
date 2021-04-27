package request

type Basic struct {
	timestamp int16 `json:"timestamp"`
	token string `json:"token"`
	data map[string]interface{} `json:"data"`
}