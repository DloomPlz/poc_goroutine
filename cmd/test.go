package cmd

import (
	"fmt"
	"poc/core"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  `test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		core.Test()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
