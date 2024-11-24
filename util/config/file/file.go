package file

import (
	"crmeb_go/util/config/parser"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	Path        string
	content     []byte
	fileExtName string
	fileLoader  *parser.Parser
}

func NewConfig(path string) *Config {
	return &Config{
		Path:       path,
		fileLoader: parser.NewFileParser(),
	}
}

// Load 读取本地文件逻辑
func (f *Config) Load() (map[string]interface{}, error) {
	f.fileExt(f.Path)
	if f.fileLoader.FileParser[f.fileExtName] == nil {
		return nil, errors.New("不支持该文件类型")
	}

	err := f.readFile()
	if err != nil {
		return nil, err
	}
	config := make(map[string]interface{})

	err = f.fileLoader.FileParser[f.fileExtName].Load(f.content, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (f *Config) fileExt(filePath string) {
	ext := filepath.Ext(filePath)
	if ext == "" {
		f.fileExtName = ext
	}
	f.fileExtName = ext[1:]
}

func (f *Config) readFile() error {
	var err error
	f.content, err = os.ReadFile(f.Path)
	return err
}
