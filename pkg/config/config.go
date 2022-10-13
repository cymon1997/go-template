package config

import "github.com/spf13/viper"

// New config from path to destination struct (pointer)
func New(dest interface{}, param Param) error {
	v := viper.New()
	v.SetConfigName(param.Name)
	v.AddConfigPath(param.Path)
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	v.AutomaticEnv()
	if param.Prefix != "" {
		v.SetEnvPrefix(param.Prefix)
	}
	err = v.Unmarshal(dest)
	if err != nil {
		return err
	}
	return nil
}
