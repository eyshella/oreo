package config

type Config struct {
	Accuracy float64
}

var AppConfig Config

func init() {
	AppConfig = Config{
		Accuracy: 0.00001,
	}
}
