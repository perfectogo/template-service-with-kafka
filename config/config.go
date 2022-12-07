package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	KafkaUrl string

	BranchServiceURL string

	CDN                 string
	MinioBucketName     string
	MinioEndpoint       string
	MinioAccessKeyID    string
	MinioSecretAccesKey string
	Bucket              string
	MinioExcelBucket    string

	LogLevel string
	HttpPort string

	DefaultLimit  string
	DefaultOffset string

	SiteCDN string

	CashierURL      string
	CashierLocalURL string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	// Postgres
	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "main"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1001"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "templates"))

	// Kafka
	config.KafkaUrl = cast.ToString(getOrReturnDefault("KAFKA_URL", "localhost:9092"))
	config.BranchServiceURL = cast.ToString(getOrReturnDefault("BRANCH_SERVICE_URL", "https://api.admin.car24.uz"))

	// Port
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8009"))
	config.DefaultOffset = cast.ToString(getOrReturnDefault("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefault("DEFAULT_LIMIT", "10"))

	config.CashierURL = cast.ToString(getOrReturnDefault("CASHIER_URL", "https://api.admin.car24.uz/v1/system-user/"))
	config.CashierLocalURL = cast.ToString(getOrReturnDefault("CASHIER_URL", "http://localhost:8001/v1/system-user/"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
