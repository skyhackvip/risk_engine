package api

import (
	"github.com/gin-gonic/gin"
	. "github.com/skyhackvip/risk_engine/api/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/run", DslRunHandler)
	return router
}
