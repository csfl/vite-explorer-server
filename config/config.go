package viteconfig

import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

func LoadConfig (env string)  {
	if env == "production" {
		config.Load(
			file.NewSource(
				file.WithPath("config/base.json"),
			),
			file.NewSource(
				file.WithPath("config/production.json"),
			),
		)
	} else {
		config.Load(
			file.NewSource(
				file.WithPath("config/base.json"),
			),
		)
	}
}