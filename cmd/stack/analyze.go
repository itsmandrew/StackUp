package stack

import (
	"fmt"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "View spending analysis based off flags and filters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Currently no implementation for analyze!")
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
