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

	db.Preload("User").Preload("Book").Find(&loans)

	for _, loan := range loans {
		if loan.Status != "dikembalikan" {
			// Menghitung selisih hari antara return_date dengan end_date menggunakan Sub()
			daysLate := int(math.Ceil(now.Sub(loan.End_date).Hours() / 24))

			if daysLate > 0 {

				// Simpan perubahan ke database
				loan.Status = "telat"
				loan.Penalty = daysLate * penalty
				if err := db.Save(&loan).Error; err != nil {
					fmt.Println("error : ", err.Error())
				}

				// Buat dan simpan notifikasi
				fullName := loan.User.Firstname + " " + loan.User.Lastname
				notification := models.Notification{
					User_id: loan.User_id,
					Message: "Buku dengan judul " + loan.Book.Title + " terlambat dikembalikan oleh " + fullName,
				}
				if err := controllers.SaveNotification(notification); err != nil {
					fmt.Println("error saving notification: ", err.Error())
				}
			}
		}
	}
}
