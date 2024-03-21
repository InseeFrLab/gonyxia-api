package configuration

import (
	_ "embed"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
)

var Config Configuration
var k = koanf.New(".")

func LoadConfiguration(defaultConfiguration string) {
	k.Load(rawbytes.Provider([]byte(defaultConfiguration)), yaml.Parser())
	k.Load(file.Provider("config.local.yaml"), yaml.Parser())
	k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil)
	k.Unmarshal("", &Config)
}
