package sedeers

import (
	"fmt"
	"pustaka-api/config"
	"pustaka-api/src/models"
)

func CreateSedeerLanguages() {
	db := config.InitConfig()
	var existingLanguages []models.Language

	if err := db.Find(&existingLanguages).Error; err != nil {
		fmt.Printf("Failed to check existing languages in the database: %s\n", err.Error())
		return
	}

	languages := []string{"Bahasa Indonesia", "Bahasa Inggris", "Bahasa Arab", "Bahasa Jepang", "Bahasa Korea", "Bahasa Mandarin", "Bahasa Perancis", "Bahasa Rusia", "Bahasa Thailand", "Bahasa Vietnam"}

	if len(existingLanguages) == 0 {
		for _, value := range languages {
			createLanguage := models.Language{
				Language: value,
			}

			if err := db.Create(&createLanguage).Error; err != nil {
				fmt.Printf("Failed to save languages %s to database: %s\n", value, err.Error())
			}
		}

		fmt.Println("==> All languages created successfully <==")
	}
}
