package cmd

import (
	"context"
	"fmt"
	"poc/core"
	"time"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  `test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		core.Test(ctx)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
