package cmd

import (
	internals "fintual-cli/internals/repositories"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var banksCmd = &cobra.Command{
	Use:   "banks",
	Short: "fintual cli is a CLI tool to query the fintual API",
	Long:  "fintual cli is a CLI tool to query the endpoints on the fintual API and track your earnings and data if you provide your account",
	Run: func(cmd *cobra.Command, args []string) {
		banks, err := internals.FintualAPIRepository.GetBanks()
		if err != nil {
			log.Fatalf("Error querying the API: %s", err)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 2, 3, ' ', 0)

		for _, bank := range banks {
			fmt.Fprintf(writer, "%s\t%s\t%s\n", bank.Id, bank.Type, bank.Attributes.Name)
		}

		writer.Flush()
	},
}

func init() {
	RootCmd.AddCommand(banksCmd)
}
