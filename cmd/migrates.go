package main

import (
	"fmt"
	"arcivum/configs"
	"arcivum/models"
)

func main() {
	configs.InitDB()

	configs.DB.Migrator().DropTable(
		&models.Codes{},
		&models.APIKeys{},
		&models.Users{},
		&models.Profiles{},
		&models.Gemini{},
		&models.GeminiHistory{},
		&models.Subscriptions{},
		&models.Avatars{},
	)

	// err := configs.DB.AutoMigrate(
	// 	&models.Codes{},
	// 	&models.APIKeys{},
	// 	&models.Users{},
	// 	&models.Profiles{},
	// 	&models.Gemini{},
	// 	&models.GeminiHistory{},
	// 	&models.Subscriptions{},
	// 	&models.Avatars{},
	// )

	// if err != nil {
	// 	fmt.Println("Migration failed:", err)
	// 	return
	// }

	fmt.Println("Migrate complete!")
}
