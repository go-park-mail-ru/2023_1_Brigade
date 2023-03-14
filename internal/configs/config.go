package configs

type Config struct {
	DB             string `yaml:"db"`
	Port           string `yaml:"port"`
	ConnectionToDB string `yaml:"connectionToDB"`

	AllowMethods     []string `yaml:"allowMethods"`
	AllowOrigins     []string `yaml:"allowOrigins"`
	AllowHeaders     []string `yaml:"allowHeaders"`
	AllowCredentials bool     `yaml:"allowCredentials"`
}
