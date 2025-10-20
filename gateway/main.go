package main

import (
	"temulokal-microservice/gateway/cmd"
	"temulokal-microservice/shared-service/logger"
)

func main() {
	logger.Init()
	cmd.Execute()
}
