package main

import (
	"github.com/VictorBion/gopportunities/config"
	"github.com/VictorBion/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main") 
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initializer()
}