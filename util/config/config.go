package config

import (
	"crmeb_go/util/config/file"
	"crmeb_go/util/config/nacos"
	"encoding/json"
	"fmt"
)

type FileReader interface {
	Load() (map[string]interface{}, error)
}

func New(c interface{}, confFile string) {
	var reader FileReader
	if confFile != "" {
		reader = file.NewConfig(confFile)
	} else {
		reader = nacos.NewConfig("yaml")
	}
	load(&c, reader)
}

func load(c interface{}, reader FileReader) {
	rawConfig, err := reader.Load()
	if err != nil {
		panic(err)
	}

	if rawConfig == nil {
		panic("config load error：must provide a config content")
	}

	configBytes, err := json.Marshal(rawConfig)
	if err != nil {
		panic(fmt.Errorf("failed to marshal config: %w", err))
	}

	if err := json.Unmarshal(configBytes, &c); err != nil {
		panic(fmt.Errorf("failed to unmarshal config into struct: %w", err))
	}
}
