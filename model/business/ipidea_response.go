package business

type IPIdeaResponse struct {
	Code           int       `mapstructure:"code" json:"code" yaml:"code"`
	Msg            string    `mapstructure:"msg" json:"msg" yaml:"msg"`
	Success        bool      `mapstructure:"success" json:"success" yaml:"success"`
	IPIdeaRespData []ProxyIP `mapstructure:"data" json:"data" yaml:"data"`
}
