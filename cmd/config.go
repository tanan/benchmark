package cmd

type AppConfig struct {
	Global struct {
		RequestCount int    `mapstructure:"request_count"`
		Url          string `mapstructure:"url"`
	} `mapstructure:"global"`
	RequestRules []struct {
		Name    string `mapstructure:"name"`
		Request struct {
			Path   string `mapstructure:"path"`
			Method string `mapstructure:"method"`
		} `mapstructure:"request"`
	} `mapstructure:"request_rules"`
}
