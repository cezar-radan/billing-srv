package cmd

import (
	"github.com/ledgertech/billing-srv/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createInvoiceCmd)
}

var createInvoiceCmd = &cobra.Command{
	Use:   "generateInvoices",
	Short: "create invoice files - pdf files, for testing purpose",
	Run: func(cmd *cobra.Command, args []string) {
		app.GenerateFakeInvoices()
	},
}
