package config

type ServerConf struct {
	System     System     `mapstructure:"system" json:"system" yaml:"system"`
	Auth       Auth       `mapstructure:"auth" json:"auth" yaml:"auth"`
	Log        Log        `mapstructure:"log" json:"log" yaml:"log"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	DB         DB         `mapstructure:"db" json:"db" yaml:"db"`
	MQ         MQ         `mapstructure:"mq" json:"mq" yaml:"mq"`
	OSS        OSS        `mapstructure:"oss" json:"oss" yaml:"oss"`
	Proxy      Proxy      `mapstructure:"proxy" json:"proxy" yaml:"proxy"`
	ThirdParty ThirdParty `mapstructure:"third-party" json:"third-party" yaml:"third-party"`
}

// System Node
type System struct {
	Env  string `mapstructure:"env-mode" json:"env-mode" yaml:"env-mode"`
	Addr int    `mapstructure:"listen-addr" json:"listen-addr" yaml:"listen-addr"`
	Db   string `mapstructure:"db" json:"db" yaml:"db"`
	Oss  string `mapstructure:"oss" json:"oss" yaml:"oss"`
}

// Auth Node
type Auth struct {
	AuthBasic AuthBasic `mapstructure:"basic" json:"basic" yaml:"basic"`
	AuthJwt   AuthJwt   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

type AuthBasic struct {
	PrivateKey string `mapstructure:"private-key" json:"private-key" yaml:"private-key"`
}

type AuthJwt struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`
	ExpiresTime int    `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"`
	BufferTime  int    `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
}

// Log Node
type Log struct {
	Director string `mapstructure:"director" json:"director" yaml:"director"`
	Level    string `mapstructure:"level" json:"level" yaml:"level"`
}

// Redis Node
type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

// DB Node
type DB struct {
	Mysql DBMysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type DBMysql struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	DBName      string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConn int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConn int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	Option      string `mapstructure:"option" json:"option" yaml:"option"`
}

// MQ Node
type MQ struct {
	RabbitMQ MqRabbit `mapstructure:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
}

type MqRabbit struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	DefaultVhost string `mapstructure:"default-vhost" json:"default-vhost" yaml:"default-vhost"`
}

// OSS Node
type OSS struct {
	Local OSSLocal `mapstructure:"local" json:"local" yaml:"local"`
}

type OSSLocal struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}

// Proxy Node
type Proxy struct {
	Url string `mapstructure:"url" json:"url" yaml:"url"`
}

// ThirdParty Node
type ThirdParty struct {
	ThirdPartyVaffle ThirdPartyVaffle `mapstructure:"vaffle" json:"vaffle" yaml:"vaffle"`
}

type ThirdPartyVaffle struct {
	InsertTransItem string `mapstructure:"call-insert-trans" json:"call-insert-trans" yaml:"call-insert-trans"`
	PrivateKey      string `mapstructure:"private-key" json:"private-key" yaml:"private-key"`
}
