package usercmd

import (
	internals "fintual-cli/internals/repositories"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var obtainTokenCmd = &cobra.Command{
	Use:   "obtain-token",
	Short: "Create a new user",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := internals.ConfigRepository.GetConfig()
		if err != nil {
			log.Fatalf("Cannot read the config file")
		}
		if config.User.Email == "" && config.User.Password == "" {
			log.Fatalf("Email or Password are not settled")
		}

		token, err := internals.FintualAPIRepository.ObtainToken(config.User.Email, config.User.Password)

		if err != nil {
			fmt.Printf("Cannot retrieve the token for error: %v", err)
		}

		internals.ConfigRepository.SetConfig("token", token, "user")
	},
}
