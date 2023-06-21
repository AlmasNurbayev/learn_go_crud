package main

import (
	"flag"
	"log"

	"github.com/AlmasNurbayev/learn_go_crud/internal/app/apiserver"
	"github.com/spf13/viper"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to config files")

}

func main() {

	flag.Parse()
	println("read config from: ", configPath)

	viper.SetConfigFile(configPath) // name of config file (without extension)
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		log.Fatal(err)
	}

	config := apiserver.NewConfig()
	config.BindAddr = viper.GetString("bind_addr")
	config.LogLevel = viper.GetString("log_level")

	viper.SetConfigFile("configs/.env") // name of config file (without extension)
	err2 := viper.ReadInConfig()        // Find and read the config file
	if err2 != nil {                    // Handle errors reading the config file
		log.Fatal(err2)
	}

	config.Store.DatabaseURL = viper.GetString("database_url")

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
