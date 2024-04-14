package config

type Database struct {
	DbDriver string `mapstructure:"DB_DRIVER"`
	DbString string `mapstructure:"DB_STRING"`
}
