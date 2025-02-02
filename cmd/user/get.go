package usercmd

import (
	internals "fintual-cli/internals/repositories"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [options]",
	Short: "Create a new user",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		emailBool, err := cmd.Flags().GetBool("email")
		if err != nil {
			log.Fatalln("Error in email flag")
		}

		passwordBool, err := cmd.Flags().GetBool("password")
		if err != nil {
			log.Fatalln("Cannot set the password")
		}

		config, err := internals.ConfigRepository.GetConfig()
		if err != nil {
			log.Fatalln("Couldn't open the config")
		}

		if emailBool && config.User.Email == "" {
			fmt.Println("Email is not settled")
		} else if emailBool {
			fmt.Printf("%s\n", config.User.Email)
		}

		if passwordBool && config.User.Password == "" {
			fmt.Println("Password is not settled")
		} else if passwordBool {
			fmt.Printf("%v\n", config.User.Password)
		}
	},
}

func init() {
	getCmd.Flags().BoolP("email", "e", false, "Set the user email")
	getCmd.Flags().BoolP("password", "p", false, "Set the user password")
}
