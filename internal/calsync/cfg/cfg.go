package cfg

import (
	"fmt"

	"github.com/spf13/viper"
)

type CalSyncConfig struct {
	From CalendarConfig `mapstructure:"from"`
	To   CalendarConfig `mapstructure:"to"`
}

type CalendarConfig struct {
	CredentialsFile string `mapstructure:"credentials_file"`
	TokensFile      string `mapstructure:"tokens_file"`
	CalendarID      string `mapstructure:"calendar_id"`
}

func ReadConfigFromFile(path string) (CalSyncConfig, error) {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return CalSyncConfig{}, fmt.Errorf("viper read calsync config error: %w", err)
	}

	cfg := CalSyncConfig{}
	if err := v.Unmarshal(&cfg); err != nil {
		return CalSyncConfig{}, fmt.Errorf("unmarshal calsync config error: %w", err)
	}

	return cfg, nil
}
