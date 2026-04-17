package cmd

import (
	"context"
	"sync"

	"github.com/Conflux-Chain/fluent-backend/api"
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/cmd"
	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func start(*cobra.Command, []string) {
	var wg sync.WaitGroup
	_, cancel := context.WithCancel(context.Background())

	// config
	var config struct {
		API     api.Config
		Service service.Config
	}
	err := viper.Unmarshal(&config)
	cmd.FatalIfErr(err, "Failed to unmarshal config")

	// services
	services, err := service.New(config.Service)
	cmd.FatalIfErr(err, "Failed to create services")

	// api
	go api.MustServe(config.API, services)

	logrus.Info("Fluent backend service started")

	cmd.GracefulShutdown(&wg, cancel)
}
