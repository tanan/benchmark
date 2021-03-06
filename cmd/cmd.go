package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"runtime"
)

var cfgFile string
var config AppConfig

var RootCmd = &cobra.Command{
	Use:   "gbench",
	Short: "gbench is a specific url benchmark tool.",
	Long: `A Fast and Simple benchmark command line tool
				written by Golang.`,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default $HOME/.gbench.yaml)")
}

func initConfig() {
	//viper.AddConfigPath("$HOME")
	//viper.SetConfigName("gbench.yaml")
	viper.SetConfigFile("/Users/toshifumi.anan/uzabase/go/src/github.com/tanan/benchmark/.gbench.yaml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newDefaultClient() (*Client, error) {
	endpointURL := config.Global.Url
	httpClient := &http.Client{}
	userAgent := fmt.Sprintf("%s (%s)", Version, runtime.Version())
	return NewClient(endpointURL, httpClient, userAgent)
}
