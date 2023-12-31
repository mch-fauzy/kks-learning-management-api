package main

//go:generate go run github.com/swaggo/swag/cmd/swag init
//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/kks-learning-management-api/configs"
	"github.com/kks-learning-management-api/shared/logger"
)

var config *configs.Config

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Initialize logger
	logger.InitLogger()

	// Initialize config
	config = configs.Get()

	// Set desired log level
	logger.SetLogLevel(config)

	// Wire everything up
	http := InitializeService()

	// Run server
	http.SetupAndServe()
}
