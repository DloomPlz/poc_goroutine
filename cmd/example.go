package cmd

import (
	"fmt"
	"poc/core"

	"github.com/spf13/cobra"
)

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "example",
	Long:  `example`,
	RunE:  run,
}

func init() {
	rootCmd.AddCommand(exampleCmd)
}

func run(cmd *cobra.Command, args []string) error {
	fmt.Println("example command called")
	return core.Example(cmd.Context())
}
