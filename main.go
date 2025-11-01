package main

import (
	"github.com/VictorBion/gopportunities/config"
	_ "github.com/VictorBion/gopportunities/docs"
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
