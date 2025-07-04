package reconciliation

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"

	"reconciliation/internal/domain"
)

type ReconciliationRepository interface {
	ParseSystemTransactions(string, time.Time, time.Time) ([]domain.Transaction, error)
	ParseBankStatements(string, time.Time, time.Time) ([]domain.BankStatement, error)
}

type reconciliationRepository struct{}

func NewReconciliationRepository() ReconciliationRepository {
	return &reconciliationRepository{}
}

func (p *reconciliationRepository) ParseSystemTransactions(filePath string, startDate, endDate time.Time) ([]domain.Transaction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions []domain.Transaction
	for _, row := range rows[1:] {
		date, err := time.Parse(time.RFC3339, row[3])
		if err != nil {
			continue
		}

		if date.Before(startDate) || date.After(endDate.Add(23*time.Hour+59*time.Minute+59*time.Second)) {
			continue
		}

		amount, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			continue
		}

		tx := domain.Transaction{
			TrxID:           row[0],
			Amount:          amount,
			Type:            domain.TransactionType(strings.ToUpper(row[2])), // normalize
			TransactionTime: date,
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func (p *reconciliationRepository) ParseBankStatements(filePath string, startDate, endDate time.Time) ([]domain.BankStatement, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var statements []domain.BankStatement
	for _, row := range rows[1:] {
		date, err := time.Parse("2006-01-02", row[2])
		if err != nil {
			continue
		}
		if date.Before(startDate) || date.After(endDate.Add(23*time.Hour+59*time.Minute+59*time.Second)) {
			continue
		}

		identifier := strings.ToUpper(row[0])
		bankName := "default_bank"
		if strings.Contains(identifier, "BCA") {
			bankName = "BCA"
		} else if strings.Contains(identifier, "MANDIRI") {
			bankName = "Mandiri"
		}

		amount, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			continue
		}

		statement := domain.BankStatement{
			UniqueIdentifier: row[0],
			Amount:           amount,
			Date:             date,
			BankName:         bankName,
		}
		statements = append(statements, statement)
	}
	return statements, nil
}
