package cmd

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	RootCmd.AddCommand(newRunCmd())
}

func newRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Manage Run resources",
		RunE:  runCmd,
	}
	return cmd
}

func runCmd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("resource name is required.")
	} else if args[0] == "help" {
		return cmd.Help()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	for _, rule := range config.RequestRules {
		httpRequest, err := client.newRequest(ctx, rule.Request.Method, rule.Request.Path, nil)
		if err != nil {
			return err
		}
		httpResponse, err := client.HTTPClient.Do(httpRequest)
		if err != nil {
			fmt.Errorf("request error: %v", err)
			continue
		}
		fmt.Printf("url: %v, method: %v, status: %v\n", config.Global.Url+rule.Request.Path, rule.Request.Method, httpResponse.Status)
	}
	return nil
}
