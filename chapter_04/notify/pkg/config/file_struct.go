package config

import "github.com/spf13/viper"

type FileConfig struct {
	Path     string
	Name     string
	FileType string
}

type Config struct {
	Vip *viper.Viper
}
