package config

type Config struct {
	Logger Logger
}

type Logger struct {
	Level    string `example:"DEBUG|INFO|WARN|ERROR"`
	Encoding string `example:"console|json"`
}

func New() Config {
	return Config{
		Logger: Logger{
			Level:    "DEBUG",
			Encoding: "console",
		},
	}
}
