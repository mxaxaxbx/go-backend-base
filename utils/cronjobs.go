package utils

import (
	"os"
	"time"

	"github.com/robfig/cron"
)

var c = cron.New()

func InitJobs() {
	defer c.Start()

	// cron to delete log files
	c.AddFunc("30 10 * * *", func() {
		Log("Running job to delete log files")
		now := time.Now().UTC().AddDate(0, 0, -5)
		date := now.Format("20060102")
		path := "logs/" + date + ".log"

		println(path)

		_, err := os.Stat(path)
		if err != nil {
			Log(err.Error())
			return
		}

		os.Remove(path)
	})

}
