package cmd

import (
	"context"
	"sync"

	"github.com/Conflux-Chain/fluent-backend/api"
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/fluent-backend/store"
	"github.com/Conflux-Chain/go-conflux-util/cmd"
	storeUtil "github.com/Conflux-Chain/go-conflux-util/store"
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
		Store   storeUtil.Config
	}
	err := viper.Unmarshal(&config)
	cmd.FatalIfErr(err, "Failed to unmarshal config")

	// store
	db := config.Store.MustOpenOrCreate(store.AllTables...)
	rawStore := storeUtil.NewStore(db)
	defer rawStore.Close()

	// services
	services, err := service.New(config.Service, rawStore)
	cmd.FatalIfErr(err, "Failed to create services")

	// api
	go api.MustServe(config.API, services)

	logrus.Info("Fluent backend service started")

	cmd.GracefulShutdown(&wg, cancel)
}
