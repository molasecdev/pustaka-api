package utils

import (
	"fmt"
	"math"
	"pustaka-api/config"
	"pustaka-api/src/models"
	"time"
)

func AutoDeleteNotification() {
	db := config.InitConfig()
	var notifications []models.Notification
	now := time.Now()

	db.Preload("User").Find(&notifications)

	for _, notif := range notifications {
		if notif.DeletedAt != nil {
			timeDeleteNotif := *notif.DeletedAt
			daysLate := int(math.Ceil(now.Sub(timeDeleteNotif).Hours() / 24))

			if notif.Read && daysLate > 0 {
				if err := db.Delete(&notif).Error; err != nil {
					fmt.Println("error : ", err.Error())
				}
			}
		}
	}
}
