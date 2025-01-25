package configs

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		App
		HTTP
		Database
	}

	App struct {
		Name    string `env-required:"" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"" yaml:"port" env:"HTTP_PORT"`
	}

	Database struct {
		DBHost     string `env-required:"" yaml:"db_host" env:"DB_HOST"`
		DBPort     string `env-required:"" yaml:"db_port" env:"DB_PORT"`
		DBPassword string `env-required:"" yaml:"db_pass" env:"DB_PASS"`
		DBName     string `env-required:"" yaml:"db_name" env:"DB_NAME"`
		DBUser     string `env-required:"" yaml:"db_user" env:"DB_USER"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	err := cleanenv.ReadConfig(configPath, config)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
