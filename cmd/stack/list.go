package stack

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists transactions based off flags and filters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Currently no implementation for list!")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
