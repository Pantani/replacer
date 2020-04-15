package cmd

import (
	"github.com/Pantani/replacer/internal/api"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "provide a replacer API",
	Run:   provideApi,
}

func provideApi(_ *cobra.Command, args []string) {
	api.Provide(Storage)
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
