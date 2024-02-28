package main

import (
	"fmt"
	"pustaka-api/config"
	"pustaka-api/sedeers"
	"pustaka-api/src/routes"
	"pustaka-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron/v2"
)

func init() {
	config.LoadEnv()
	config.InitConfig()
}

func cronjob() {
	schedule, errSchedule := gocron.NewScheduler()
	if errSchedule != nil {
		fmt.Println("errSchedule : ", errSchedule)
	}

	_, errjob := schedule.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				// NewAtTime(hours,minutes,seconds) format 24 jam
				gocron.NewAtTime(0, 1, 0),
				gocron.NewAtTime(10, 0, 0),
				gocron.NewAtTime(14, 0, 0),
				gocron.NewAtTime(18, 0, 0),
			),
		),
		gocron.NewTask(
			func() {
				utils.AutoUpdateLateStatusAndPenalty()
				utils.AutoDeleteNotification()
			},
		),
	)
	if errjob != nil {
		fmt.Println("errjob : ", errjob)
	}

	schedule.Start()
}

func main() {
	router := gin.Default()
	// routes
	v1 := router.Group("/v1")
	routes.InitRoutes(v1)

	// sedeers
	sedeers.InitSedeers()

	// automatic running on 23:00 everyday
	cronjob()

	router.Run()
}
