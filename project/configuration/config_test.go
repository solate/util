package configuration

import (
	"testing"

	config2 "github.com/solate/util/project/configuration/app/config"
)

func TestLoadConfig(t *testing.T) {

	cfg := &config2.TomlConfig{}
	err := LoadConfig(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(cfg)
}
