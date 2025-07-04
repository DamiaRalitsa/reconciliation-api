package http

import (
	usecases "reconciliation/internal/usecases/reconciliation"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ReconciliationController struct {
	useCase usecases.ReconciliationUsecase
}

func NewReconciliationController() *ReconciliationController {
	return &ReconciliationController{
		useCase: usecases.NewReconciliationUsecase(),
	}
}

func (h *ReconciliationController) Reconcile(c *fiber.Ctx) error {
	var req struct {
		SystemFilePath string   `json:"system_file_path"`
		BankFilePaths  []string `json:"bank_file_paths"`
		StartDate      string   `json:"start_date"`
		EndDate        string   `json:"end_date"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if req.SystemFilePath == "" {
		req.SystemFilePath = "assets/templates/system_transaction.csv"
	}
	if len(req.BankFilePaths) == 0 {
		req.BankFilePaths = []string{
			"assets/templates/bank_bca_statements.csv",
			"assets/templates/bank_mandiri_statements.csv",
		}
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid start date format"})
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid end date format"})
	}

	summary, err := h.useCase.Reconcile(req.SystemFilePath, req.BankFilePaths, startDate, endDate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(summary)
}
