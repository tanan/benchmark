package cmd

type AppConfig struct {
	Global       Global        `mapstructure:"global"`
	RequestRules []RequestRule `mapstructure:"request_rules"`
}

type Global struct {
	Url      string `mapstructure:"url"`
	Count    int    `mapstructure:"count"`
	Parallel int    `mapstructure:"parallel"`
}

type RequestRule struct {
	Name    string  `mapstructure:"name"`
	Count   int     `mapstructure:"count"`
	Request Request `mapstructure:"request"`
}

type Request struct {
	Path   string `mapstructure:"path"`
	Method string `mapstructure:"method"`
}
