package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func SetConfig() {

	viper.SetConfigName("conf") // name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath("./libs/config/") // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	os.Setenv("ENVIRONMENT", "DEVELOPMENT")
	viper.AutomaticEnv()

	fmt.Println("ENV ", viper.GetString("database.mongodb.host"), viper.Get("ENVIRONMENT"))
}
