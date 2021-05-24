package config

type BasicConfig struct {
	Nacos Nacos `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
}

type Nacos struct {
	Namespace   string `mapstructure:"namespace" json:"namespace" yaml:"namespace"`
	LogDir      string `mapstructure:"log-dir" json:"log-dir" yaml:"log-dir"`
	CacheDir    string `mapstructure:"cache-dir" json:"cache-dir" yaml:"cache-dir"`
	IpAddr      string `mapstructure:"ip-addr" json:"ip-addr" yaml:"ip-addr"`
	ContextPath string `mapstructure:"ctx-path" json:"ctx-path" yaml:"ctx-path"`
	Port        uint64 `mapstructure:"port" json:"port" yaml:"port"`
	Scheme      string `mapstructure:"scheme" json:"scheme" yaml:"scheme"`
	DataId      string `mapstructure:"data-id" json:"data-id" yaml:"data-id"`
	Group       string `mapstructure:"group" json:"group" yaml:"group"`
}
