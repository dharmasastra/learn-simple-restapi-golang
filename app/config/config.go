package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect		string
	Location	string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:"sqlite3",
			Location:"data.db",
		},
	}
}
