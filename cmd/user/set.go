package usercmd

import (
	internals "fintual-cli/internals/repositories"
	"log"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [options]",
	Short: "Create a new user",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			log.Fatalln("Error in email flag")
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalln("Cannot set the password")
		}

		if email != "" {
			err = internals.ConfigRepository.SetConfig("email", email)
			if err != nil {
				log.Printf("Cannot set the email for %v\n", err)
			}
		}

		if password != "" {
			err = internals.ConfigRepository.SetConfig("password", password)
			if err != nil {
				log.Printf("Cannot set the password for %v\n", err)
			}
		}
	},
}

func init() {
	setCmd.Flags().StringP("email", "e", "", "Set the user email")
	setCmd.Flags().StringP("password", "p", "", "Set the user password")
}
