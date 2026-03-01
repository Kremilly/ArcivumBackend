package main

import (
	"fmt"
	"arcivum/configs"
	"arcivum/models"
)

func main() {
	configs.InitDB()

	configs.DB.Migrator().DropTable(
		// &models.APIKeys{},
		// &models.Backups{},
		// &models.Users{},
		// &models.Profiles{},
		// &models.Databases{},
		// &models.Gemini{},
		// &models.GeminiHistory{},
		// &models.Subscriptions{},
	)

	err := configs.DB.AutoMigrate(
		&models.Codes{},
		&models.APIKeys{},
		&models.Backups{},
		&models.Databases{},
		&models.Users{},
		&models.Profiles{},
		&models.Gemini{},
		&models.GeminiHistory{},
		&models.Subscriptions{},
		&models.Avatars{},
		&models.DumpAccessLogs{},
		&models.TeamMembers{},
		&models.Teams{},
	)

	if err != nil {
		fmt.Println("Migration failed:", err)
		return
	}

	fmt.Println("Migrate complete!")
}
