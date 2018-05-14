package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Clients - struct for client threading configuration
type Clients struct {
	Count uint32
	Fiber uint32
}

// Data - struct for data source configuration
type Data struct {
	FileName   string
	Duration   []uint16
	PacketSize uint16 `yaml:"packet-size"`
}

// Config - struct for application configuration
type Config struct {
	Clients
	Data
}

// New create of Cfg struct and try to load it from yaml file
func New(name string) (*Config, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}
	viper.SetConfigName(name)
	viper.AddConfigPath(dir)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	cfg := Config{
		Clients: Clients{
			Count: 2,
			Fiber: 4,
		},
		Data: Data{
			Duration:   []uint16{750, 1000, 3000},
			PacketSize: 100,
		},
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
