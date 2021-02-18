package configreader

import "github.com/spf13/viper"

// NewConfigReader returns a CongiReader instance
func NewConfigReader(name, configType, path string) (*ConfigReader, error) {
	viper.SetConfigName(name)
	viper.SetConfigType(configType)
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return &ConfigReader{}, nil
}

// ConfigReader is a struct to read config file with help of viper package
type ConfigReader struct{}

// GetString return some value as string from config file or env based on key
func (cr *ConfigReader) GetString(key string) string {
	return viper.GetString(key)
}
