package main

import (
	"temulokal-microservice/auth-service/cmd"
	"temulokal-microservice/shared-service/logger"
)

func main() {
	logger.Init()
	cmd.Execute()
}
