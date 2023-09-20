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
}

type DatabaseGroup struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type SmtpConfig struct {
	SmtpServer  string `yaml:"smtp_server"`
	SmtpPort    int    `yaml:"smtp_port"`
	SenderEmail string `yaml:"sender_email"`
	Password    string `yaml:"password"`
}

type ServerConfig struct {
	SecretKey string `yaml:"secret_key"`
}

type Configs struct {
	UIDir         string        `yaml:"ui_dir"`
	Database      DatabaseGroup `yaml:"database"`
	MRExpiryTime  int64         `yaml:"mr_expiry_time"`
	TokenLifeTime int32         `yaml:"token_life_time"`
	Smtp          SmtpConfig    `yaml:"smtp_config"`
	Server        ServerConfig  `yaml:"server"`
	Redis         RedisConfig   `yaml:"redis_connection"`
}
