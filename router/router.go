package router

import (
	"github.com/gin-gonic/gin"
)

func Initializer() {

	r := gin.Default()

	initializerRoutes(r)
	
	r.Run()
}