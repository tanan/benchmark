package cmd

type AppConfig struct {
	Global struct {
		Count int    `mapstructure:"count"`
		Url   string `mapstructure:"url"`
	} `mapstructure:"global"`
	RequestRules []struct {
		Name    string `mapstructure:"name"`
		Count   int    `mapstructure:"count"`
		Request struct {
			Path   string `mapstructure:"path"`
			Method string `mapstructure:"method"`
		} `mapstructure:"request"`
	} `mapstructure:"request_rules"`
}
