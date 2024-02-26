package utils

import (
	"fmt"
	"math"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"time"
)

func UpdateLateStatusAndPenalty() {
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
				loan.Status = "telat"
				loan.Penalty = daysLate * penalty

				// Simpan perubahan ke database
				if err := db.Save(&loan).Error; err != nil {
					fmt.Println("error : ", err.Error())
				}
			}
		}
	}
}
