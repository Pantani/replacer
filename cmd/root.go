package cmd

import (
	"fmt"
	"github.com/Pantani/logger"
	"github.com/Pantani/replacer/internal/config"
	"github.com/Pantani/replacer/internal/storage"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var (
	Storage *storage.Storage
	rootCmd = &cobra.Command{
		Use:   "Backend Interview",
		Short: "Interview test for backend developer",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger.Info("Connecting to the database...")
			var err error
			Storage, err = storage.New()
			if err != nil {
				logger.Fatal(err)
			}
			logger.Info("Database connected")
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	loadConf := func() { config.InitConfig() }
	cobra.OnInitialize(loadConf)
}
