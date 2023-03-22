package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	apppkg "github.com/perfectogo/template-service-with-kafka/internal/app"
	configpkg "github.com/perfectogo/template-service-with-kafka/internal/pkg/config"
)

var roodCmd = &cobra.Command{
	Use:   "grcp-server",
	Short: "To run grpc server do not give arguments",
	Run: func(cmd *cobra.Command, args []string) {
		// initialization config
		config := configpkg.New()

		// initialization app
		app, err := apppkg.NewApp(config)
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			if err := app.Run(); err != nil {
				app.Logger.Error("app run", zap.Error(err))
			}
		}()

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs

		app.Logger.Info("Payment app service stops")

		// app stops
		app.Stop()
	},
}

func Execute() {
	if err := roodCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
