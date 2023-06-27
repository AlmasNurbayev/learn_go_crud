package apiserver

type Config struct {
	BindAddr    string
	LogLevel    string
	DatabaseURL string
	LogPath     string
	KeyJwt      string
	// Store    *sqlstore.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "8083",
		LogLevel: "debug",
		KeyJwt:   "aszc;lkypewv",
		// Store:    sqlstore.NewConfig(),
	}
}
