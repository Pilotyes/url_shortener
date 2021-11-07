package config

// MainConfig is main config struct
type MainConfig struct {
	BindAddr string `toml:"bind_addr"`
}

// NewConfig returns new MainConfig
func NewConfig() *MainConfig {
	return &MainConfig{
		BindAddr: ":8000",
	}
}
