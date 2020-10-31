package cmd

import (
	"fmt"
	"poc/core"

	"github.com/spf13/cobra"
)

var limiterCmd = &cobra.Command{
	Use:   "limiter",
	Short: "limiter",
	Long:  `limiter`,
	RunE:  runLimiter,
}

func init() {
	rootCmd.AddCommand(limiterCmd)
}

func runLimiter(cmd *cobra.Command, args []string) error {
	fmt.Println("limiter command called")
	return core.Limiter(cmd.Context())
}
