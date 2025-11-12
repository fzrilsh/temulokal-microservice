package main

import (
	"temulokal-microservice/shared-service/logger"
	"temulokal-microservice/umkm-service/cmd"
)

func main() {
	logger.Init()
	cmd.Execute()
}
