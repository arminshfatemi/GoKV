package config

type Config struct {
	Server struct {
		Address string `json:"address"`
	} `json:"server"`

	Auth struct {
		Enabled  bool   `json:"enabled"`
		SaltFile string `json:"salt_file"`
	} `json:"auth"`

	Users []UserConfig `json:"users"`
}

type UserConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (c *Config) applyDefaults() {
	if c.Server.Address == "" {
		c.Server.Address = ":6379"
	}
}

func (c *Config) Validate() error {
	return nil
}
