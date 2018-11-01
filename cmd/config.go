package cmd

type AppConfig struct {
	Global       GlobalConfig
	RequestRules []RequestRuleConfig
}

type GlobalConfig struct {
	RequestCount int    `yaml:"request_count"`
	Url          string `yaml:"url"`
}

type RequestRuleConfig struct {
	Name    string  `yaml:"name"`
	Request Request `yaml:"request"`
}

type Request struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}
