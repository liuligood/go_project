package iconfig

import (
	"crmeb_go/utils/iconfig/file"
	"crmeb_go/utils/iconfig/nacos"
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
		panic("iconfig load error：must provide a iconfig content")
	}

	configByteList, err := json.Marshal(rawConfig)
	if err != nil {
		panic(fmt.Errorf("failed to marshal iconfig: %w", err))
	}

	if err := json.Unmarshal(configByteList, &c); err != nil {
		panic(fmt.Errorf("failed to unmarshal iconfig into struct: %w", err))
	}

}
