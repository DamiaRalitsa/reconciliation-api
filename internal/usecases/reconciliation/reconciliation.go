package reconciliation

import (
	"time"

	"reconciliation/internal/domain"
	"reconciliation/internal/repositories/reconciliation"
)

type ReconciliationUsecase interface {
	Reconcile(systemFilePath string, bankFilePaths []string, startDate, endDate time.Time) (domain.ReconciliationDetails, error)
}

type reconciliationUsecase struct {
	reconciliationRepo reconciliation.ReconciliationRepository
}

func NewReconciliationUsecase() ReconciliationUsecase {
	reconciliationRepo := reconciliation.NewReconciliationRepository()
	return &reconciliationUsecase{reconciliationRepo: reconciliationRepo}
}

func (s *reconciliationUsecase) Reconcile(systemFilePath string, bankFilePaths []string, startDate, endDate time.Time) (domain.ReconciliationDetails, error) {
	systemTx, err := s.reconciliationRepo.ParseSystemTransactions(systemFilePath, startDate, endDate)
	if err != nil {
		return domain.ReconciliationDetails{}, err
	}

	var bankStatements []domain.BankStatement
	for _, file := range bankFilePaths {
		bs, err := s.reconciliationRepo.ParseBankStatements(file, startDate, endDate)
		if err != nil {
			return domain.ReconciliationDetails{}, err
		}
		bankStatements = append(bankStatements, bs...)
	}

	summary := MatchTransactions(systemTx, bankStatements)
	return summary, nil
}
