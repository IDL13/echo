package utils

import (
	"log"
	"os"
)

func Loger(err error) {
	file, _ := os.Create("log.txt")
	logger := log.New(file, "INFO/t", log.Ldate|log.Ltime)
	logger.Println(err)
	defer file.Close()
}
