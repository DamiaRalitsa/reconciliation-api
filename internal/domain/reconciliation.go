package domain

import "time"

type TransactionType string

const (
	Debit  TransactionType = "DEBIT"
	Credit TransactionType = "CREDIT"
)

type Transaction struct {
	TrxID           string          `json:"trx_id"`
	Amount          float64         `json:"amount"`
	Type            TransactionType `json:"type"`
	TransactionTime time.Time       `json:"transaction_time"`
}

type BankStatement struct {
	UniqueIdentifier string    `json:"unique_identifier"`
	Amount           float64   `json:"amount"`
	Date             time.Time `json:"date"`
	BankName         string    `json:"bank_name"`
}

type ReconciliationDetails struct {
	TotalTransactions           int                        `json:"total_transactions"`
	MatchedTransactions         int                        `json:"matched_transactions"`
	UnmatchedSystemTransactions []Transaction              `json:"unmatched_system_transactions"`
	UnmatchedBankStatements     map[string][]BankStatement `json:"unmatched_bank_statements"`
	TotalDiscrepanciesAmount    float64                    `json:"total_discrepancies_amount"`
}
