package app

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/perfectogo/template-service-with-kafka/internal/infrastructure/grpc_service_clients"
	postgresql_repo "github.com/perfectogo/template-service-with-kafka/internal/infrastructure/repository/postgresql"
	"github.com/perfectogo/template-service-with-kafka/internal/pkg/config"
	"github.com/perfectogo/template-service-with-kafka/internal/pkg/logger"
	"github.com/perfectogo/template-service-with-kafka/internal/pkg/otlp"
	"github.com/perfectogo/template-service-with-kafka/internal/pkg/postgres"
	investUsecase "github.com/perfectogo/template-service-with-kafka/internal/usecase"
	"github.com/perfectogo/template-service-with-kafka/internal/usecase/event"
)

type ScheduleCron struct {
	Config         *config.Config
	Logger         *zap.Logger
	DB             *postgres.PostgresDB
	BrokerProducer event.BrokerProducer
	ShutdownOTLP   func() error
	ServiceClients grpc_service_clients.ServiceClients
}

func NewScheduleCron(cfg *config.Config) (*ScheduleCron, error) {
	// initialization logger
	logger, err := logger.New(cfg.LogLevel, cfg.Environment, cfg.APP+".log")
	if err != nil {
		return nil, err
	}

	// otlp collector initialization
	shutdownOTLP, err := otlp.InitOTLPProvider(cfg)
	if err != nil {
		return nil, err
	}

	db, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ScheduleCron{
		Config:       cfg,
		Logger:       logger,
		DB:           db,
		ShutdownOTLP: shutdownOTLP,
	}, nil
}

func (a *ScheduleCron) Run(ctx context.Context) error {
	// context timeout initialization
	contextTimeout, err := time.ParseDuration(a.Config.Context.Timeout)
	if err != nil {
		return fmt.Errorf("error during parse duration for context timeout : %w", err)
	}
	// Initialize Service Clients
	serviceClients, err := grpc_service_clients.New(a.Config)
	if err != nil {
		return fmt.Errorf("error during initialize service clients: %w", err)
	}
	a.ServiceClients = serviceClients

	// repositories initialization
	investmentPostgresRepo := postgresql_repo.NewInvestmentRepo(a.DB)
	invesmtnetTrRepo := postgresql_repo.NewInvestmentTransactionRepo(a.DB)
	invesmtnetPortfolioRepo := postgresql_repo.NewInvestmentPortfolioRepo(a.DB)
	dividendTransactionRepo := postgresql_repo.NewDividendTransactionRepo(a.DB)
	partnershipValueRepo := postgresql_repo.NewPartnershipValuesRepo(a.DB)
	withdrawalSettingRepo := postgresql_repo.NewWithdrawalSettingsRepo(a.DB)
	outgoingTransactionRepo := postgresql_repo.NewOutgoingTransactionRepo(a.DB)
	paymetnTransactionPostgresRepo := postgresql_repo.NewPaymentTransactionRepo(a.DB)
	groupConditionRepo := postgresql_repo.NewGroupConditionsRepo(a.DB)
	tariffRepo := postgresql_repo.NewTariffRepo(a.DB)
	cronStatusesRepo := postgresql_repo.NewCronStatusesRepo(a.DB)

	tariffUseCase := investUsecase.NewTariff(contextTimeout, tariffRepo, a.BrokerProducer, a.Logger)
	groupConditionUseCase := investUsecase.NewGroupCondition(contextTimeout, groupConditionRepo)
	outgoingTransactionUseCase := investUsecase.NewOutgoingPaymentTransaction(contextTimeout, outgoingTransactionRepo, a.BrokerProducer, a.Logger)
	investmentUseCase := investUsecase.NewInvestment(contextTimeout, investmentPostgresRepo, a.BrokerProducer, a.Logger)
	invesmtnetTrUseCase := investUsecase.NewInvestmentTransaction(contextTimeout, invesmtnetTrRepo, a.BrokerProducer, a.Logger)
	invesmtnetPortfolioUseCase := investUsecase.NewInvestmentPortfolio(contextTimeout, investmentUseCase, invesmtnetTrUseCase, invesmtnetPortfolioRepo, a.BrokerProducer, a.Logger)
	paymentTransactionUseCase := investUsecase.NewPaymentTransaction(contextTimeout, paymetnTransactionPostgresRepo, a.Logger, a.BrokerProducer)
	withdrawalSettingsUseCase := investUsecase.NewWithdrawalSetting(contextTimeout, withdrawalSettingRepo, a.BrokerProducer, a.Logger)
	partnershipValueUseCase := investUsecase.NewPartnershipValues(
		contextTimeout,
		a.Logger,
		partnershipValueRepo,
		investmentUseCase,
		invesmtnetTrUseCase,
		paymentTransactionUseCase,
		groupConditionUseCase,
		tariffUseCase,
		a.BrokerProducer,
	)
	cronStatusesUseCase := investUsecase.NewCronStatuses(contextTimeout, cronStatusesRepo)

	dividendTransactionUseCase := investUsecase.NewDividendTransaction(
		contextTimeout,
		dividendTransactionRepo,
		withdrawalSettingsUseCase,
		investmentUseCase,
		outgoingTransactionUseCase,
		invesmtnetTrUseCase,
		paymentTransactionUseCase,
		partnershipValueUseCase,
		invesmtnetPortfolioUseCase,
		withdrawalSettingsUseCase,
		cronStatusesUseCase,
		a.Logger,
		a.BrokerProducer,
	)

	timeNow := time.Now()
	// Capitalize cron
	err = dividendTransactionUseCase.Capitalize(ctx)
	if err != nil {
		return fmt.Errorf("error during Capitalize: %w", err)
	}

	a.Logger.Info("Scheduled work completed",
		zap.String("process_duration", time.Since(timeNow).String()))
	return nil
}

func (a *ScheduleCron) Stop() {
	// close broker producer
	if a.BrokerProducer != nil {
		a.BrokerProducer.Close()
	}
	// closing client service connections
	if a.ServiceClients != nil {
		a.ServiceClients.Close()
	}
	// database connection
	a.DB.Close()
	// shutdown otlp collector
	if err := a.ShutdownOTLP(); err != nil {
		a.Logger.Error("shutdown otlp collector", zap.Error(err))
	}
	// zap logger sync
	a.Logger.Sync()
}
