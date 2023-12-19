package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	logger.Info("application is starting")
	app.Start()
}
