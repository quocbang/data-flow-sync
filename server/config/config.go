package config

// Options for the implementation
type Options struct {
	ConfigPath string `short:"c" long:"config" description:"server configuration file" required:"true"`
}

type PostgresConfig struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type DatabaseGroup struct {
	Postgres PostgresConfig `yaml:"postgres"`
	Redis    RedisConfig    `yaml:"redis"`
}

type Configs struct {
	UIDir         string              `yaml:"ui-dir"`
	Database      DatabaseGroup       `yaml:"database"`
	AccessDir     []string            `yaml:"access-dir"`
	MRExpiryTime  int64               `yaml:"mr-expiry-time"`
	TokenLifeTime int32               `yaml:"token-life-time"`
	Roles         map[string][]string `yaml:"roles"`
}
