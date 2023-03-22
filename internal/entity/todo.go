package entity

import "time"

const (
	// statuses
	StatusNew        = "new"
	StatusProcessed  = "processed"
	StatusSuccess    = "success"
	StatusPending    = "pending"
	StatusWaiting    = "waiting"
	StatusActive     = "active"
	StatusVerified   = "verified"
	StatusInvestment = "investment"
	StatusActual     = "actual"
	StatusClosed     = "closed"
	StatusEnable     = "enable"
	StatusDisable    = "disable"
	StatusCancelled  = "cancelled"
	StatusChecked    = "checked"
	StatusApproved   = "approved"
	StatusCompleted  = "completed"
	StatusRejected   = "rejected"

	// transaction types
	TransactionTypeInvestment = "investment"
	TransactionTypeOutgoing   = "outgoing"
	TransactionTypePayment    = "payment"
	TransactionTypeDividend   = "dividend"

	// calculation methods
	CalculationMethodPlus  = "plus"
	CalculationMethodMinus = "minus"

	// operation ids
	OperationAccrualProfit          int64 = 1
	OperationDeductionIrr           int64 = 2
	OperationDeductionTax           int64 = 3
	OperationPartnershipShare       int64 = 4
	OperationPartnershipProfit      int64 = 5
	OperationWithdrawalDividends    int64 = 6
	OperationRefundInvestment       int64 = 7
	OperationDividendCapitalization int64 = 8
	OperationInvestment             int64 = 9
	OperationCancelTransaction      int64 = 10
	OperationWithdrawalInvestment   int64 = 11

	// payment type
	PaymentTypeCard = "card"
	PaymentTypeBank = "bank"

	// pl_values
	IrrValue = "irr_value"

	// group conditions
	MaxValue         = "max_value"
	StartValue       = "start_value"
	StepBalance      = "step_balance"
	StepTransaction  = "step_transaction"
	StepActivity     = "step_activity"
	DeltaBalance     = "delta_balance"
	DeltaTransaction = "delta_transaction"
	DeltaActivity    = "delta_activity"
	MaxBalance       = "max_balance"
	MaxTransaction   = "max_transaction"
	MaxActivity      = "max_activity"
	StepTerm         = "step_term"
	DeltaTerm        = "delta_term"
	MaxTerm          = "max_term"
	MinBalance       = "min_balance"

	// calculation values
	TaxNoresident      = "tax_noresident"
	TaxResident        = "tax_resident"
	TaxCorporate       = "tax_corporate"
	AutoApproveValue   = "autoapprove_value"
	ExitLimit          = "exit_limit"
	MinInvestmentSaldo = "min_investment_saldo"

	// Cron status investment
	CronCapitalizationStatus = "capitalization"
	CronInvestStatus         = "invest"
	CronPortfolioStatus      = "portfolio"
	CronPartnershipStatus    = "partnership"

	// group names
	GroupNameOld = "old"
)

// operation types
var OperationTypes = map[int64]string{
	1:  "accrual_profit",
	2:  "deduction_irr",
	3:  "deduction_tax",
	4:  "partnership_share",
	5:  "partnership_profit",
	6:  "withdrawal_dividends",
	7:  "refund_investment",
	8:  "dividend_capitalization",
	9:  "investment",
	10: "cancel_transaction",
	11: "withdrawal_investment",
}

type Investment struct {
	GUID           string
	InvestorID     string
	Status         string
	TariffID       string
	GroupID        string
	GoalIconID     string
	StrategyID     string
	GoalTitle      string
	ReferalId      string
	InvestorStatus string
	IsResident     bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type InvestmentNetProfit struct {
	GUID         string
	InvestmentID string
	Month        time.Time
	Value        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CronStatus struct {
	ID        int64
	Status    string
	Stage     string
	Month     time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
