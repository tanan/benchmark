package cmd

import (
	"context"
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

	for _, rule := range config.RequestRules {
		res, execTime := runBench(client, rule, config.Global.Parallel)
		res.RenderOutput(execTime)
	}

	return nil
}

func runBench(client *Client, rule RequestRule, parallel int) (ResponseResultList, time.Duration) {

	var result []ResponseResult

	httpStream := make(chan bool, parallel)
	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < rule.Count; i++ {
		wg.Add(1)
		httpStream <- true
		go func() error {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			httpRequest, err := client.newRequest(ctx, rule.Request.Method, rule.Request.Path, rule.Request.Header, nil)
			if err != nil {
				return err
			}
			s := time.Now()
			httpResponse, err := client.HTTPClient.Do(httpRequest)
			if err != nil {
				return err
			}
			res := ResponseResult{
				Url:       httpRequest.URL.Host + httpRequest.URL.Path,
				Method:    httpRequest.Method,
				Status:    httpResponse.StatusCode,
				TimeTaken: time.Since(s),
			}

			result = append(result, res)
			<-httpStream
			return nil
		}()
	}
	wg.Wait()
	return result, time.Since(start)
}
