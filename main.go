package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"arcivum/configs"

	// "arcivum/controllers/auth"
	// "arcivum/controllers/dumps"
	// "arcivum/controllers/dumps/tools"
	// "arcivum/controllers/dumps/vault"
	// "arcivum/controllers/gemini"
	// "arcivum/controllers/gemini/history"
	// "arcivum/controllers/keys"
	// "arcivum/controllers/plans"
	// "arcivum/controllers/profile"
	// "arcivum/controllers/teams"
	// "arcivum/controllers/databases"

	// "arcivum/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/helmet"
)

func main() {
	configs.InitDB()
	configs.ConnectRedis()

	if configs.Redis == nil {
		log.Println("WARNING: Redis was not initialized. Cache will not work.")
	}

	app := fiber.New(fiber.Config{
		AppName:     configs.ProductName + " Services v1",
		TrustProxy:  true,
		ProxyHeader: "CF-Connection-IP",
	})

	app.Use(helmet.New())

	app.Use(func(c fiber.Ctx) error {
		raw := c.Get("Cookie")
		if raw != "" {
			parts := strings.Split(raw, ";")
			valid := make([]string, 0, len(parts))

			for _, p := range parts {
				p = strings.TrimSpace(p)
				if strings.Contains(p, "=") && len(p) > 10 {
					valid = append(valid, p)
				}
			}

			if len(valid) > 0 {
				c.Request().Header.Set("Cookie", strings.Join(valid, "; "))
			} else {
				c.Request().Header.Del("Cookie")
			}
		}
		return c.Next()
	})

	app.Use(func(c fiber.Ctx) error {
		origin := c.Get("Origin")
		if origin == "" {
			origin = "*"
		}

		c.Set("Access-Control-Allow-Origin", origin)
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Csrf-Token")
		c.Set("Access-Control-Allow-Credentials", "true")

		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(http.StatusNoContent)
		}

		return c.Next()
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Welcome to Arcivum Services v1")
	})

	// app.Get("/auth/me", auth.GetProfile)
	// app.Delete("/auth/logout", auth.Logoff)
	// app.Post("/auth/login", middleware.LoginLimiter(), auth.LoginUser)
	// app.Get("/auth/has-logged", auth.HasLogged)
	// app.Post("/auth/register", auth.CreateUser)

	// app.Get("/databases", databases.ListDatabases)
	// app.Get("/databases/:id", databases.GetDatabase)
	// app.Get("/databases/:id/export", databases.GetDumpSettings)

	// app.Get("/backups/:id", dumps.ListDumps)
	// app.Post("/backups/create", dumps.CreateDump)

	// app.Get("/backups/:user/:id", dumps.GetDump)
	// app.Get("/backups/:user/:id/raw", dumps.GetDumpRaw)
	// app.Get("/backups/:user/:id/parsed", dumps.GetDumpParser)
	// app.Put("/backups/:user/:id/privacy", dumps.SetDumpPrivacy)
	// app.Get("/backups/:user/:id/structure", dumps.GetDumpStructure)

	// app.Post("/backups/:user/:id/xss-scan", tools.ScanDumpForXSS)
	// app.Post("/backups/:user/:id/export-csv", tools.ExportToCSV)
	// app.Post("/backups/:user/:id/convert-model", tools.ConvertSQLToModel)

	// app.Get("/backups/:user/:id/access", vault.ListAccessLogs)
	// app.Post("/backups/:user/:id/unlock", vault.CreateAccessLogs)

	// app.Post("/teams/create", teams.CreateTeam)

	// app.Get("/profile/:user", profile.GetProfile)
	// app.Get("/profile/:user/check", profile.CheckUsername)
	// app.Get("/profile/:user/dumps", profile.ListProfileDumps)

	// app.Get("/plans/list", plans.PlansList)

	// app.Get("/gemini/history", history.ListGeminiHistory)
	// app.Get("/gemini/history/:id", history.GetGeminiItemHistory)
	// app.Delete("/gemini/history/clear", history.ClearGeminiHistory)

	// app.Get("/settings/keys", keys.GetAPIKeys)
	// app.Put("/settings/keys/:slug/renew", keys.RenewApiKey)
	// app.Put("/settings/keys/:slug/status", keys.UpdateStatusApi)
	// app.Put("/settings/keys/:slug", keys.UpdatePermissionsApiKey)

	// app.Put("/settings/profile", profile.EditProfile)
	// app.Put("/settings/profile/username", profile.EditUsername)

	// app.Put("/settings/gemini", gemini.EditGemini)
	// app.Get("/settings/gemini", gemini.GetSettings)
	// app.Get("/settings/gemini/models", gemini.ListAvailableModels)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
