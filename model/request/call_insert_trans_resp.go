package request

type CallInsertTransResp struct {
	Code int         `mapstructure:"code" json:"code" validate:"code"`
	Data interface{} `mapstructure:"data" json:"data" validate:"data"`
	Msg  string      `mapstructure:"msg" json:"msg" validate:"msg"`
}
