package utils

import (
	xerrors "github.com/pkg/errors"
	"github.com/spf13/viper"
)

func GetConfig(path string, name string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(name)
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		return nil, xerrors.WithMessage(err, "读取配置文件错误")
	}
	return v, nil
}
