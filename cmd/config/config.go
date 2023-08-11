package config

import "github.com/spf13/viper"

type Config struct {
	HttpPort  string   `mapstructure:"HTTP_PORT"`
	SecretKey string   `mapstructure:"SECRET_KEY"`
	Database  Database `mapstructure:",squash"`
	Casbin    Casbin   `mapstructure:",squash"`
}

type Database struct {
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	Name     string `mapstructure:"DB_NAME"`
}

type Casbin struct {
	ConfPath string `mapstructure:"CAS_CONF_PATH"`
	CSVPath  string `mapstructure:"CAS_CSV_PATH"`
}

func Load() (Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.ReadInConfig()

	viper.BindEnv("HTTP_PORT")
	viper.BindEnv("SECRET_KEY")
	viper.BindEnv("DB_USERNAME")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("CAS_CONF_PATH")
	viper.BindEnv("CAS_CSV_PATH")

	var c Config
	err := viper.Unmarshal(&c)

	return c, err
}
