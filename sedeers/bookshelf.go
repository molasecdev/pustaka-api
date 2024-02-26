package sedeers

import (
	"fmt"
	"pustaka-api/config"
	"pustaka-api/src/models"
)

func CreateSedeerBookshelfs() {
	db := config.InitConfig()
	var existingBookshelves []models.Bookshelfs

	if err := db.Find(&existingBookshelves).Error; err != nil {
		fmt.Printf("Failed to check existing bookshelves in the database: %s\n", err.Error())
		return
	}

	if len(existingBookshelves) == 0 {
		for i := 'A'; i <= 'D'; i++ {
			bookcase := string(i)
			for j := 1; j <= 4; j++ {
				for k := 1; k <= 6; k++ {
					col := fmt.Sprintf("%d-%d", j, k)
					no_bookshelf := bookcase + col

					bookshelf := models.Bookshelfs{
						Bookshelf: no_bookshelf,
					}
					if err := db.Create(&bookshelf).Error; err != nil {
						fmt.Printf("Failed to save bookshelf %s to database: %s\n", no_bookshelf, err.Error())
					}
				}
			}
		}

		fmt.Println("==> All bookshelves created successfully <==")
	}
}
