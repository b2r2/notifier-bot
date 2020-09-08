package app

// Config ...
type Config struct {
	Token       string   `toml:"token"`
	ChatID      int64    `toml:"chat_id"`
	LogLevel    string   `toml:"log_level"`
	BotLogLevel bool     `toml:"bot_log_level"`
	AccessUsers []string `toml:"access_users"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		LogLevel:    "debug",
		BotLogLevel: true,
	}
}
