package stack

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Spins up a dashboard visualization in the browser for the user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Currently there is no implementation for the dashboard command")
	},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
