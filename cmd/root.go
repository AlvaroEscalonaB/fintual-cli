package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fintual",
	Short: "fintual cli is a CLI tool to query the fintual API",
	Long:  "fintual cli is a CLI tool to query the endpoints on the fintual API and track your earnings and data if you provide your account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to fintual CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
