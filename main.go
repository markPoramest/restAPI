package main

import (
	"restAPI/app"
	"restAPI/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
