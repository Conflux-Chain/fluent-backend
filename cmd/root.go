package cmd

import (
	"github.com/Conflux-Chain/go-conflux-util/config"
	"github.com/Conflux-Chain/go-conflux-util/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "fluent-backend",
	Short: "Start backend service for Fluent wallet",
	Run:   start,
}

func init() {
	cobra.OnInitialize(config.MustInitDefault)

	log.BindFlags(&rootCmd)
}

// Execute is the command line entrypoint.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Failed to execute command")
	}
}
