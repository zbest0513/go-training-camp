package data

import (
	"github.com/google/wire"
	"notify-server/internal/data/ent"
	"notify-server/internal/pkg"
	"notify/pkg/config"
	"notify/pkg/utils"
)

var ProviderSet = wire.NewSet(NewTagRepo, NewUserRepo, NewTemplateRepo)

var SqlClientProviderSet = wire.NewSet(NewEntClient, NewConfig)

func NewEntClient(vip *config.Config) *ent.Client {
	return pkg.NewClient(vip)
}

func NewConfig(fc config.FileConfig) (*config.Config, error) {
	vip, err := utils.GetConfig(fc.Path, fc.Name, fc.FileType)
	//config, err := utils.GetConfig("./configs", "config", "yaml")
	if err != nil {
		return nil, err
	}
	return &config.Config{
		vip,
	}, nil
}
