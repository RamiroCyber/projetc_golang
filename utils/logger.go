package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Logger(prefix, message string) {
	currentTime := time.Now().Format("2006-01-02T15-04-05")
	filename := fmt.Sprintf("%s-%s.log", prefix, currentTime)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	customLogger := log.New(file, "", log.LstdFlags)
	customLogger.Println(message)

}
