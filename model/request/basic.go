package request

type Basic struct {
	Timestamp int                    `mapstructure:"Timestamp" json:"timestamp" validate:"required"`
	Token     string                 `mapstructure:"Token" json:"token" validate:"required"`
	Data      map[string]interface{} `mapstructure:"data" json:"data" validate:"required"`
}
