package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"net/http"
)

func init() {
	global.Features = dto.NewGlobalFeatures()
	global.DslResult = dto.NewDslResult()
}

func DslRunHandler(c *gin.Context) {
	var request dto.DslRunRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	for k, v := range request.Features {
		f := dto.Feature{Name: k, Value: v}
		global.Features.Set(f)
	}

	flow := request.Flow
	dsl := dslparser.LoadDslFromFile("test/yaml/" + flow + ".yaml")
	rs := dsl.Parse(global.DslResult)
	c.JSON(http.StatusOK, gin.H{
		"flow":   flow,
		"result": rs,
	})
}
