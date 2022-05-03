package cmd

import (
	"github.com/ledgertech/billing-srv/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts server",
	Run: func(cmd *cobra.Command, args []string) {
		app.Start()
	},
}
