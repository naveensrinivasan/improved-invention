package tests

type Config struct {
	Db struct {
		Name       string `yaml:"name"`
		User       string `yaml:"user"`
		Password   string `yaml:"password"`
		Port       int    `yaml:"port"`
		Initialize bool   `yaml:"initialize"`
	} `yaml:"db"`
	Standard struct {
		TimeoutInSeconds           int `yaml:"timeoutinseconds"`
		PollIntervalInMilliseconds int `yaml:"pollintervalinmilliseconds"`
	} `yaml:"standard"`
}
