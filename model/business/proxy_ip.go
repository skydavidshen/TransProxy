package business

type ProxyIP struct {
	IP   string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}
