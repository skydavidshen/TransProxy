package request

type Header struct {
	Timestamp []string `mapstructure:"Timestamp" validate:"required"`
	Token []string `mapstructure:"Token" validate:"required"`
	ContentType []string `mapstructure:"Content-Type"`
	ContentLength []string `mapstructure:"Content-Length"`
	UserAgent []string `mapstructure:"User-Agent"`
}
