package usercmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Query and set the user information",
	Long:  "Query and set the user information",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add email or password to see your goals")
	},
}

func init() {
	UserCmd.AddCommand(getCmd)
	UserCmd.AddCommand(setCmd)
}
