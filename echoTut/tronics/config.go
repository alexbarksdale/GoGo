package tronics

// ConfigDatabase holds configuration properties based on env variables.
type ConfigDatabase struct {
	AppName  string `env:"APP_NAME" env-default:"TRONICS"`
	AppEnv   string `env:"APP_ENV" env-default:"DEV"`
	Port     string `env:"PORT" env-default:"8080"`
	Host     string `env:"HOST" env-default:"localhost"`
	LogLevel string `env:"LOG_LEVEL" env-default:"ERROR"`
}

var cfg ConfigDatabase
