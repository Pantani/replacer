package config

import (
	"github.com/Pantani/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type configuration struct {
	Port  string
	Mongo struct {
		Uri string
	}
	Api struct {
		Mode string
	}
}

var Configuration configuration

// set dummy values to force viper to search for these keys in environment variables
// the AutomaticEnv() only searches for already defined keys in a config file, default values or kvstore struct.
func setDefaults() {
	viper.SetDefault("Port", "3000")
	viper.SetDefault("Mongo.Uri", "mongodb://replacer:MongoDB2019!@mongo-database:27017")
	viper.SetDefault("Api.Mode", "release")
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	setDefaults()
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.Unmarshal(&Configuration); err != nil {
		logger.Error(err, "Error Unmarshal Viper Config File")
	}
	log.Printf("API_PORT: %s", Configuration.Port)
	log.Printf("MONGO_URI: %s", Configuration.Mongo.Uri)
	log.Printf("API_MODE: %s", Configuration.Api.Mode)
}
