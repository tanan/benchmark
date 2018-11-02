package cmd

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"sync"
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

	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	httpStream := make(chan bool, 10)
	var wg sync.WaitGroup

	for _, rule := range config.RequestRules {
		for i := 0; i < rule.Count; i++ {
			wg.Add(1)
			httpStream <- true
			go func() error {
				defer wg.Done()
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				httpRequest, err := client.newRequest(ctx, rule.Request.Method, rule.Request.Path, nil)
				if err != nil {
					return err
				}
				start := time.Now()
				httpResponse, err := client.HTTPClient.Do(httpRequest)
				if err != nil {
					return err
				}
				defer httpResponse.Body.Close()
				fmt.Printf("url: %v, method: %v, status: %v, took_time: %v\n", httpRequest.URL, rule.Request.Method, httpResponse.Status, time.Since(start))
				<-httpStream
				return nil
			}()
		}
	}
	wg.Wait()
	return nil
}
