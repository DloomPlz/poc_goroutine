package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"poc/core"
	"syscall"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  `test`,
	RunE:  run,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func run(cmd *cobra.Command, args []string) error {
	fmt.Println("test called")
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		signal.Stop(sigs)
		cancel()
	}()
	go func() {
		select {
		case <-sigs:
			cancel()
		case <-ctx.Done():
		}
	}()
	return core.Test(ctx)
}

// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
// 	done := make(chan bool, 1)
// 	core.Test(done)
// 	go func() {
// 		sig := <-sigs
// 		fmt.Println()
// 		fmt.Println(sig)
// 		done <- true
// 	}()
// 	<-done
// 	return nil
