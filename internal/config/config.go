package config

import (
	"github.com/spf13/viper"
)

// Config yapısı, YAML dosyasındaki hiyerarşi ile birebir aynı olmalı
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type DatabaseConfig struct {
	URL            string `mapstructure:"url"`
	MaxOpenConns   int    `mapstructure:"max_open_conns"`
	MaxIdleConns   int    `mapstructure:"max_idle_conns"`
	MigrationsPath string `mapstructure:"migrations_path"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")  // Projenin ana klasörüne bak
	viper.AddConfigPath("..") // cmd/ altından çalıştırılınca bir üst dizine bak

	// Çevresel değişkenleri (ENV) oku (Örn: DATABASE_URL varsa YAML'dakini ezer)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
