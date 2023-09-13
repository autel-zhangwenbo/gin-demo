package configs

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {


	// default read toml type config file from workdir
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	parseConfig()

	log.Infof("global cfg.Mqtt: %+v", Cfg.Mqtt)
}


func parseConfig() {
	once.Do(func() {
		err := viper.Unmarshal(&Cfg)
		if err != nil {
			log.Fatalln(err)
		}
	})
}