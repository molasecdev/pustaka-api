package utils

import (
	"fmt"
	"math"
	"pustaka-api/config"
	"pustaka-api/src/controllers"
	"pustaka-api/src/models"
	"time"
)

func AutoUpdateLateStatusAndPenalty() {
	db := config.InitConfig()
	var loans []models.Loan

	penalty := 2000
	now := time.Now()

	db.Find(&loans)

	for _, loan := range loans {
		if loan.Status != "dikembalikan" {
			// Menghitung selisih hari antara return_date dengan end_date menggunakan Sub()
			daysLate := int(math.Ceil(now.Sub(loan.End_date).Hours() / 24))

			if daysLate > 0 {
				var user models.User
				var book models.Book

				if err := db.Where("id = ?", loan.User_id).First(&user).Error; err != nil {
					fmt.Println("User with ID", loan.User_id, "not found")
					continue
				}
				if err := db.Where("id = ?", loan.Book_id).First(&book).Error; err != nil {
					fmt.Println("Book with ID", loan.Book_id, "not found")
					continue
				}

				// Simpan perubahan ke database
				loan.Status = "telat"
				loan.Penalty = daysLate * penalty
				if err := db.Save(&loan).Error; err != nil {
					fmt.Println("error : ", err.Error())
				}

				// Buat dan simpan notifikasi
				fullName := user.Firstname + " " + user.Lastname
				notification := models.Notification{
					User_id: user.ID,
					Message: "Buku dengan judul " + book.Title + " terlambat dikembalikan oleh " + fullName,
				}
				if err := controllers.SaveNotification(notification); err != nil {
					fmt.Println("error saving notification: ", err.Error())
				}
			}
		}
	}
}
