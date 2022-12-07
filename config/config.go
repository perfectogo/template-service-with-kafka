package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"

	SaleCategoryID        = 1
	PurchaseCategoryID    = 2
	ExpenditureCategoryID = 3

	PaymeAccountNumber   = 10000032
	ClickAccountNumber   = 2
	CompanyAccountNumber = 10000000

	RentCategoryID                = 1
	CarExpenditureCategoryID      = 2
	EmployeeExpenditureCategoryID = 3
	TransferCategoryID            = 4
	TopUpCategoryID               = 5

	TransfeCashboxToCompanySubCategoryID = "278c8f5f-9627-4aa7-bfff-064d9dbf475b"
	TariffSubCategoryID                  = "4b7322e4-980f-4db9-82b1-8f395efd87d5"
	FuelSubCategoryID                    = "825f1720-54a8-44b2-b79c-40f1735d3ebf"
	DepositSubCategoryID                 = "993224db-d25f-43cb-bfea-a8ee842f34b8"
	FineSubCategoryID                    = "7b89250e-75d8-46a6-8e8f-844e31f1d1b7"
	RepairSubCategoryID                  = "7cefe920-afb5-4980-a2ad-275720a94a35"
	ReturnDepositSubCategoryID           = "2b9064ab-09b8-414d-adb3-a8ee1a3b7863"
	GiveCarSubCategoryID                 = "616e9497-a224-49d2-aa65-a1b1a7310858"
	CostFromDepositSubCategoryID         = "85e12a9d-5dc4-4c20-8d4d-a6015a19ed81"
	TopUpSubCategoryID                   = "79cf8a5f-416f-4e19-bf23-7132571c1f31"

	DateTimeLayout = "2006-01-02T15:04:05.999Z"
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
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "billing_service"))

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
