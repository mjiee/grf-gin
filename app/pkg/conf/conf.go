package conf

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App   `mapstructure:"app" json:"app" yaml:"app"`
	Log   `mapstructure:"log" json:"log" yaml:"log"`
	Db    `mapstructure:"db" json:"db" yaml:"db"`
	Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Jwt   `mapstructure:"jwt" jsob:"jwt" yaml:"jwt"`
}

func NewConfig(confFile string) (*Config, error) {
	viper.SetConfigFile(confFile)
	viper.ReadInConfig()

	viper.SetEnvPrefix("ADMIN")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var conf Config

	err := viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.New("无法解析项目配置")
	}

	return &conf, nil
}
