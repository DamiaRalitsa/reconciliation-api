package reconciliation

import (
	"math"
	"reconciliation/internal/domain"
)

func GroupBankStatements(banks []domain.BankStatement) map[string][]domain.BankStatement {
	bankMap := make(map[string][]domain.BankStatement)

	for _, b := range banks {
		dateKey := b.Date.Format("2006-01-02")
		bankMap[dateKey] = append(bankMap[dateKey], b)
	}

	return bankMap
}

func FindAndRemoveMatch(tx domain.Transaction, bankStatements []domain.BankStatement) (matched bool, updatedStatements []domain.BankStatement, discrepancy float64) {
	normalizedAmount := tx.Amount
	if tx.Type == domain.Debit {
		normalizedAmount = -normalizedAmount
	}

	for i, stmt := range bankStatements {
		diff := math.Abs(normalizedAmount - stmt.Amount)
		if diff <= 5000 {
			return true, append(bankStatements[:i], bankStatements[i+1:]...), diff
		}
	}

	return false, bankStatements, 0
}

func MatchTransactions(system []domain.Transaction, banks []domain.BankStatement) domain.ReconciliationDetails {
	summary := domain.ReconciliationDetails{
		UnmatchedBankStatements: make(map[string][]domain.BankStatement),
		TotalTransactions:       len(system),
	}

	bankMap := GroupBankStatements(banks)

	for _, sysTx := range system {
		dateKey := sysTx.TransactionTime.Format("2006-01-02")
		bankStmts, exists := bankMap[dateKey]

		if !exists || len(bankStmts) == 0 {
			summary.UnmatchedSystemTransactions = append(summary.UnmatchedSystemTransactions, sysTx)
			continue
		}

		matched, remainingStatements, discrepancy := FindAndRemoveMatch(sysTx, bankStmts)
		if matched {
			summary.MatchedTransactions++
			summary.TotalDiscrepanciesAmount += discrepancy
			bankMap[dateKey] = remainingStatements
		} else {
			summary.UnmatchedSystemTransactions = append(summary.UnmatchedSystemTransactions, sysTx)
		}
	}

	for _, remaining := range bankMap {
		for _, b := range remaining {
			summary.UnmatchedBankStatements[b.BankName] = append(summary.UnmatchedBankStatements[b.BankName], b)
		}
	}

	return summary
}
