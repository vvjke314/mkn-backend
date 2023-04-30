package config

type Config struct {
	DataBase DbConfig `toml:"database"`
}

type DbConfig struct {
	Host     string `toml:"db_host"`
	Name     string `toml:"db_name"`
	User     string `toml:"db_user"`
	Password string `toml:"db_password"`
	Port     string `toml:"db_port"`
}
