package common

import (
	"encoding/json"
	"os"

	"github.com/Meduzz/devserver/config"
)

func LoadFileContent(file string) ([]byte, error) {
	return os.ReadFile(file)
}

func StoreFileContent(file string, content []byte) error {
	return os.WriteFile(file, content, 0x755)
}

func LoadAppConfig(ctx *AppContext) (*config.Config, error) {
	content, err := LoadFileContent(ctx.File)

	if err != nil {
		return nil, err
	}

	cfg := &config.Config{}
	err = json.Unmarshal(content, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func StoreAppConfig(ctx *AppContext, cfg *config.Config) error {
	bs, err := json.Marshal(cfg)

	if err != nil {
		return err
	}

	return StoreFileContent(ctx.File, bs)
}

// TODO implement
func LoadProperty(file, property string) ([]byte, error) {
	return nil, nil
}

// TODO implement
func StoreProperty(file, property string, content []byte) error {
	return nil
}
