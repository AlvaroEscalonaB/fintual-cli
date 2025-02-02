package cmd

import (
	usercmd "fintual-cli/cmd/user"
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "fintual",
	Short: "fintual cli is a CLI tool to query the fintual API",
	Long:  "fintual cli is a CLI tool to query the endpoints on the fintual API and track your earnings and data if you provide your account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to fintual CLI")
	},
}

func init() {
	RootCmd.AddCommand(usercmd.UserCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
