package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type clientsCfg struct {
	count  int
	fibers int
}

type config struct {
	clients clientsCfg
}

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	viper.SetConfigName("cls")
	viper.AddConfigPath(dir)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Config file not found...")
	} else {
		cfg := config{clients: {count: 7, fibers: 7}}
		fmt.Println("Default values is:", cfg)
		err = viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatal("Unable to decode into struct", err)
		} else {
			fmt.Println("Cfg is:", cfg)
		}
	}
}
