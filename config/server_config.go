package config

type ServerConf struct {
	System        System        `mapstructure:"system" json:"system" yaml:"system"`
	Auth          Auth          `mapstructure:"auth" json:"auth" yaml:"auth"`
	Log           Log           `mapstructure:"log" json:"log" yaml:"log"`
	Redis         Redis         `mapstructure:"redis" json:"redis" yaml:"redis"`
	DB            DB            `mapstructure:"db" json:"db" yaml:"db"`
	MQ            MQ            `mapstructure:"mq" json:"mq" yaml:"mq"`
	OSS           OSS           `mapstructure:"oss" json:"oss" yaml:"oss"`
	ThirdParty    ThirdParty    `mapstructure:"third-party" json:"third-party" yaml:"third-party"`
	TransPlatform TransPlatform `mapstructure:"trans-platform" json:"trans-platform" yaml:"trans-platform"`
	Handler       Handler       `mapstructure:"handler" json:"handler" yaml:"handler"`
	Switch        Switch        `mapstructure:"switch" json:"switch" yaml:"switch"`
}

// System Node
type System struct {
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
	Username     string         `mapstructure:"username" json:"username" yaml:"username"`
	Password     string         `mapstructure:"password" json:"password" yaml:"password"`
	Addr         string         `mapstructure:"addr" json:"addr" yaml:"addr"`
	DefaultVhost string         `mapstructure:"default-vhost" json:"default-vhost" yaml:"default-vhost"`
	Option       MqRabbitOption `mapstructure:"option" json:"option" yaml:"option"`
	Expiration   string         `mapstructure:"expiration" json:"expiration" yaml:"expiration"`
}

type MqRabbitOption struct {
	Exchange MqRabbitOptionExchange `mapstructure:"exchange" json:"exchange" yaml:"exchange"`
	Queue    MqRabbitOptionQueue    `mapstructure:"queue" json:"queue" yaml:"queue"`
}

type MqRabbitOptionExchange struct {
	InsertTransItems     string `mapstructure:"insert-trans-items" json:"insert-trans-items" yaml:"insert-trans-items"`
	TransItems           string `mapstructure:"trans-items" json:"trans-items" yaml:"trans-items"`
	DeadTransItems       string `mapstructure:"dead-trans-items" json:"dead-trans-items" yaml:"dead-trans-items"`
	DeadInsertTransItems string `mapstructure:"dead-insert-trans-items" json:"dead-insert-trans-items" yaml:"dead-insert-trans-items"`
}

type MqRabbitOptionQueue struct {
	InsertTransItem     MqRabbitOptionQueueItem `mapstructure:"insert-trans-item" json:"insert-trans-item" yaml:"insert-trans-item"`
	TransItem           MqRabbitOptionQueueItem `mapstructure:"trans-item" json:"trans-item" yaml:"trans-item"`
	DeadTransItem       MqRabbitOptionQueueItem `mapstructure:"dead-trans-item" json:"dead-trans-item" yaml:"dead-trans-item"`
	DeadInsertTransItem MqRabbitOptionQueueItem `mapstructure:"dead-insert-trans-item" json:"dead-insert-trans-item" yaml:"dead-insert-trans-item"`
}

type MqRabbitOptionQueueItem struct {
	Name  string                    `mapstructure:"name" json:"name" yaml:"name"`
	Binds []MqRabbitOptionQueueBind `mapstructure:"bind" json:"bind" yaml:"bind"`
}

type MqRabbitOptionQueueBind struct {
	Exchange string `mapstructure:"exchange" json:"exchange" yaml:"exchange"`
	Key      string `mapstructure:"key" json:"key" yaml:"key"`
}

// OSS Node
type OSS struct {
	Local OSSLocal `mapstructure:"local" json:"local" yaml:"local"`
}

type OSSLocal struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}

// ThirdParty Node
type ThirdParty struct {
	ThirdPartyVaffle ThirdPartyVaffle `mapstructure:"vaffle" json:"vaffle" yaml:"vaffle"`
}

type ThirdPartyVaffle struct {
	InsertTransItem string `mapstructure:"insert-trans-item" json:"insert-trans-item" yaml:"insert-trans-item"`
	PrivateKey      string `mapstructure:"private-key" json:"private-key" yaml:"private-key"`
}

// TransPlatform Node
type TransPlatform struct {
	SmartProxy TransPlatformSmartProxy `mapstructure:"smartproxy" json:"smartproxy" yaml:"smartproxy"`
	IpIdea     TransPlatformIpIdea     `mapstructure:"ipidea" json:"ipidea" yaml:"ipidea"`
}

type TransPlatformSmartProxy struct {
	Username  string `mapstructure:"username" json:"username" yaml:"username"`
	Password  string `mapstructure:"password" json:"password" yaml:"password"`
	ProxyHost string `mapstructure:"proxy-host" json:"proxy-host" yaml:"proxy-host"`
}

type TransPlatformIpIdea struct {
	Url string `mapstructure:"url" json:"url" yaml:"url"`
}

// Handler Node
type Handler struct {
	TransItemGoroutineCount           int `mapstructure:"transitem-goroutine-count" json:"transitem-goroutine-count" yaml:"transitem-goroutine-count"`
	CallInsertTransItemGoroutineCount int `mapstructure:"call-insert-transitem-goroutine-count" json:"call-insert-transitem-goroutine-count" yaml:"call-insert-transitem-goroutine-count"`
}

// Switch Node
type Switch struct {
	AuthBasic        bool `mapstructure:"auth-basic" json:"auth-basic" yaml:"auth-basic"`
	UseRealTranslate bool `mapstructure:"use-real-translate" json:"use-real-translate" yaml:"use-real-translate"`
}
