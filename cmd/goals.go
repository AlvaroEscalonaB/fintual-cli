package cmd

import (
	internals "fintual-cli/internals/repositories"
	"fintual-cli/internals/utils"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var goalsCmd = &cobra.Command{
	Use:   "goals",
	Short: "fintual cli is a CLI tool to query the fintual API",
	Long:  "fintual cli is a CLI tool to query the endpoints on the fintual API and track your earnings and data if you provide your account",
	Run: func(cmd *cobra.Command, args []string) {
		userPayload, err := internals.ConfigRepository.GetUserPayload()
		if err != nil {
			log.Fatalf("Error getting the user payload from config: %s", err)
		}

		goals, err := internals.FintualAPIRepository.ObtainGoals(userPayload)
		if err != nil {
			log.Fatalf("Error querying the API: %s", err)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 2, 3, ' ', 0)

		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n", "Nombre", "Depositado", "Total", "Profit")

		for _, goal := range goals {
			fmt.Fprintf(writer, "%s\t%f\t%f\t%f\n", utils.CleanText(goal.Attributes.Name), goal.Attributes.Deposited, goal.Attributes.Nav, goal.Attributes.Profit)
		}

		writer.Flush()
	},
}

func init() {
	RootCmd.AddCommand(goalsCmd)
}
