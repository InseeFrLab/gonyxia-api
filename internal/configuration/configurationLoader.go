package configuration

import (
	_ "embed"
	"strings"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var s string
var Config Configuration

func LoadConfiguration() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yaml")
	viper.ReadConfig(strings.NewReader(s)) // Reading defaults
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.SetConfigName("config.local")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.MergeInConfig()

	err := viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
