package stack

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a short message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from StackUp!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
