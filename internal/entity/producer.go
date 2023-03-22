package entity

import "time"

type BonusValuesForProducer struct {
	GUID             string    `json:"guid"`
	GroupID          string    `json:"group_id"`
	BonusBalance     float64   `json:"bonus_balance"`
	BonusTransaction float64   `json:"bonus_transaction"`
	BonusActivity    float64   `json:"bonus_activity"`
	DeltaActivity    float64   `json:"delta_activity"`
	DeltaBalance     float64   `json:"delta_balance"`
	DeltaTransaction float64   `json:"delta_transaction"`
	MaxBalance       float64   `json:"max_balance"`
	MaxTransaction   float64   `json:"max_transaction"`
	MaxActivity      float64   `json:"max_activity"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type TransactionHistory struct {
	ID              int64     `json:"id"`
	GUID            string    `json:"guid"`
	TransactionID   string    `json:"transaction_id"`
	InvestmentID    string    `json:"investment_id"`
	Month           string    `json:"month"`
	Operation       string    `json:"operation"`
	TransactionType string    `json:"transaction_type"`
	PaymentType     string    `json:"payment_type"`
	PaymentSystem   string    `json:"payment_system"`
	Amount          float64   `json:"amount"`
	EndAmount       float64   `json:"end_amount"`
	Balance         float64   `json:"balance"`
	Status          string    `json:"status"`
	Date            time.Time `json:"date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CloseMonth struct {
	BonusValues          []*BonusValuesForProducer `json:"bonus_values"`
	TransactionHistories []*TransactionHistory     `json:"transaction_histories"`
}

type Aggregate struct {
	BonusValues          []*BonusValuesForProducer `json:"bonus_values"`
	TransactionHistories []*TransactionHistory     `json:"transaction_histories"`
	Investments          []*Investment             `json:"investments"`
}
