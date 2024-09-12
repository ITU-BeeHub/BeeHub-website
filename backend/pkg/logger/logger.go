package logger

import (
	"log"
)

func Info(message string) {
	log.Println("INFO: " + message)
}

func Error(err error) {
	log.Println("ERROR: " + err.Error())
}
