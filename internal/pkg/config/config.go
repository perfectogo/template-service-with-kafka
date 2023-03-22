package config

import (
	"os"
	"strings"
)

// App
type APP struct {
	App         string
	Environment string
	LogLevel    string
	RPCPort     string
}

// Context
type Context struct {
	Timeout string
}

// DB
type DB struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Sslmode  string
}

// Kafka
type Kafka struct {
	Address []string
	Topic   struct {
		InvestmentPaymentTransaction         string
		InvestmentCreated                    string
		UpdatePaymentTransactionStatusActive string
		Aggregate                            string
		PaymentWithdrawalTransaction         string
		PaymentWithdrawalTransactionStatus   string
	}
}

// OTP
type OTP struct {
	Expiration string
	StaticCode string
}

// OTLPCollector
type OTLPCollector struct {
	Host string
	Port string
}

// All
type Config struct {
	APP           APP
	Context       Context
	DB            DB
	Kafka         Kafka
	OTP           OTP
	OTLPCollector OTLPCollector
}

func New() *Config {
	var config = Config{}

	// initialization app
	config.APP.App = getEnv("APP", "app")
	config.APP.Environment = getEnv("ENVIRONMENT", "develop")
	config.APP.LogLevel = getEnv("LOG_LEVEL", "debug")
	config.APP.RPCPort = getEnv("RPC_PORT", ":50051")

	// initialization db
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "1000m")

	// initialization db
	config.DB.Host = getEnv("POSTGRES_HOST", "localhost")
	config.DB.Port = getEnv("POSTGRES_PORT", "5432")
	config.DB.Name = getEnv("POSTGRES_DATABASE", "postgres")
	config.DB.User = getEnv("POSTGRES_USER", "postgres")
	config.DB.Password = getEnv("POSTGRES_PASSWORD", "postgres")
	config.DB.Sslmode = getEnv("POSTGRES_SSLMODE", "disable")

	// initialization otp
	config.OTP.Expiration = getEnv("OTP_EXPIRATION", "30m")
	config.OTP.StaticCode = getEnv("OTP_STATIC_CODE", "00000")

	// initialization otlp collector
	config.OTLPCollector.Host = getEnv("OTLP_COLLECTOR_HOST", "localhost")
	config.OTLPCollector.Port = getEnv("OTLP_COLLECTOR_PORT", ":4317")

	// kafka address
	config.Kafka.Address = strings.Split(getEnv("KAFKA_ADDRESS", "127.0.0.1:29092"), ",")

	// kafka topics
	config.Kafka.Topic.InvestmentPaymentTransaction = getEnv("KAFKA_TOPIC_INVESTMENT_PAYMENT_TRANSACTION", "investment.payment.transaction.created")
	config.Kafka.Topic.UpdatePaymentTransactionStatusActive = getEnv("KAFKA_TOPIC_INVESTOR_PAYMENT_TRANSACTION_STATUS", "investor.payment.transaction.active")
	config.Kafka.Topic.InvestmentCreated = getEnv("KAFKA_TOPIC_INVESTMENT_CREATED", "investment.investment.created")
	config.Kafka.Topic.Aggregate = getEnv("KAFKA_TOPIC_AGGREGATE", "aggregate")
	config.Kafka.Topic.PaymentWithdrawalTransaction = getEnv("KAFKA_TOPIC_PAYMENT_WITHDRAWAL_TRANSACTION", "payment.withdrawal.transaction")
	config.Kafka.Topic.PaymentWithdrawalTransactionStatus = getEnv("KAFKA_TOPIC_PAYMENT_WITHDRAWAL_TRANSACTION_STATUS", "payment.withdrawal.transaction.status")

	return &config
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
