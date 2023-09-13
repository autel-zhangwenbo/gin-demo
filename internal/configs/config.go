package configs

import "sync"

type Config struct {
   Mqtt  *MqttConfig `mapstructure:"mqtt"`
   AppCfg  *AppConfig `mapstructure:"app"`
   Database *Database `mapstructure:"database"`
}


type MqttConfig struct {
	Node string `mapstructure:"node"`
	Url  string `mapstructure:"url"`
}

type AppConfig struct {
	Env string `mapstructure:"env"`
	AppName  string `mapstructure:"app_name"`
	HttpHost  string `mapstructure:"http_host"`
	HttpPort  string `mapstructure:"http_port"`
	EnableDocs  string `mapstructure:"enable_docs"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
	Port     int32 `mapstructure:"port"`
	ConnectionMax     int32 `mapstructure:"connection_max"`
	Enabled     string `mapstructure:"enabled"`
}



var (
	Cfg    Config
	once   sync.Once
)