package config

type System struct {
	Env           string `mapstructure:"env_mode" json:"env_mode" yaml:"env_mode"`
	Addr          int    `mapstructure:"listen_addr" json:"listen_addr" yaml:"listen_addr"`
	Db            string `mapstructure:"db" json:"db" yaml:"db"`
	Oss           string `mapstructure:"oss" json:"oss" yaml:"oss"`
}
