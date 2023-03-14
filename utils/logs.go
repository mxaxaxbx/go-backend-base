package utils

import (
	"log"
	"os"
	"time"
)

func Log(str string) {
	now := time.Now().UTC()
	date := now.Format("20060102")
	path := "logs/" + date + ".log"

	datetime := now.Format("2006-01-02 15:04:05")
	content := datetime + " - " + str

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString(content + "\n")
	log.Println(content)
}
