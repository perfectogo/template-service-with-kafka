package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	apppkg "github.com/perfectogo/template-service-with-kafka/internal/app"
	configpkg "github.com/perfectogo/template-service-with-kafka/internal/pkg/config"
)

const (
	PAYMENT_TRANSACTION_CONSUMER                    = "payment_transaction"
	PAYMENT_WITHDRAWAL_TRANSACTIONS_STATUS_CONSUMER = "payment_withdrawal_transactions_status"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "To run consumer give the name followed by arguments consumer",
	Long: `Example : 
	go run cmd/main.go consumer name_of_consumer`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		consumerName := args[0]

		switch consumerName {
		case PAYMENT_TRANSACTION_CONSUMER:
			PaymentTransactionConsumerRun()
		case PAYMENT_WITHDRAWAL_TRANSACTIONS_STATUS_CONSUMER:
			PaymentWithrdrawalTransactionsRun()
		default:
			log.Fatalf("No consumer with name: %s", consumerName)
		}

	},
}

func init() {
	roodCmd.AddCommand(consumerCmd)
}

func PaymentTransactionConsumerRun() {
	// initialization config
	config := configpkg.New()

	// initialization app
	app, err := apppkg.NewPaymentTransactionConsumer(config)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := app.Run(); err != nil {
			app.Logger.Error("payment transaction consumer run", zap.Error(err))
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	app.Logger.Info("payment transaction consumer stops")

	// app stops
	app.Stop()
}

func PaymentWithrdrawalTransactionsRun() {
	// initialization config
	config := configpkg.New()

	// initialization app
	app, err := apppkg.NewWithdrawalTransactionsConsumer(config)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := app.Run(); err != nil {
			app.Logger.Error("payment withdrawal transaction consumer run", zap.Error(err))
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	app.Logger.Info("payment withdrawal transaction consumer stops")

	// app stops
	app.Stop()
}
