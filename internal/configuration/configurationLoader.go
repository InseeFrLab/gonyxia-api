package configuration

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var Config Configuration
var k = koanf.New(".")

func LoadConfiguration() {
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Load YAML config and merge into the previously loaded config (because we can).
	k.Load(file.Provider("config.local.yaml"), yaml.Parser())
	k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil)
	k.Unmarshal("", &Config)
	fmt.Print(Config.OIDC.IssuerURI)
}
