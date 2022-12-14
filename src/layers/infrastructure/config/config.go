package config

type Config struct {
	Log    `yaml:"log"`
	Listen struct {
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

type Log struct {
	ActiveLevels []string `yaml:"active_levels" env-default:"all"`
	Output       string   `yaml:"output" env-default:"stdout"`
	PathToFile   string   `yaml:"path_to_file" env-default:"all.log"`
}
