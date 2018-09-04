package cmd

import (
	"os"
	"os/signal"

	"github.com/7phs/coding-challenge-search/config"
	"github.com/7phs/coding-challenge-search/db"
	"github.com/7phs/coding-challenge-search/logger"
	"github.com/7phs/coding-challenge-search/model"
	"github.com/7phs/coding-challenge-search/restapi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Init()

		log.Info(ApplicationInfo())

		config.Init()

		db.Init(config.Conf.DatabaseUrl)

		model.Init()

		srv := restapi.
			NewServer(config.Conf).
			Run()

		stop := make(chan os.Signal)
		signal.Notify(stop, os.Interrupt)
		<-stop
		log.Info("interrupt signal")

		srv.Shutdown()

		db.Shutdown()

		log.Info("finish")
	},
}
