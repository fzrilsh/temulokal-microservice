package main

import (
	"temulokal-microservice/umkm-service/cmd"
	"temulokal-microservice/shared-service/logger"
)

func main() {
	logger.Init()
	cmd.Execute()
}
