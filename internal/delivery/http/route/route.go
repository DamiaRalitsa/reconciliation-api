package route

import (
	"encoding/json"
	"os"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"

	"reconciliation/internal/delivery/http"
)

type RouteConfig struct {
	App                      *fiber.App
	reconciliationController *http.ReconciliationController
}

func NewRouteConfig() *RouteConfig {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	app := fiber.New(fiber.Config{
		Prefork:     false,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   100 * 1024 * 1024,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	reconciliationController := http.NewReconciliationController()

	routeConfig := RouteConfig{
		App:                      app,
		reconciliationController: reconciliationController,
	}

	routeConfig.SetupReconciliation()
	return &routeConfig
}

func (rc *RouteConfig) SetupReconciliation() {
	reconciliationGroup := rc.App.Group("/api/reconciliation")
	reconciliationGroup.Post("", rc.reconciliationController.Reconcile)
}

func (rc *RouteConfig) Listen(address string) {
	rc.App.Listen(address)
}
