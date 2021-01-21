package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"net/http"
	//"strconv"
	"fmt"
)

func init() {
	fmt.Println("init")
	global.Features = dto.NewGlobalFeatures()
	global.DslResult = dto.NewDslResult()
}

func DslRunHandler(c *gin.Context) {
	var request dto.DslRunRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println(err)
	}
	//global.Features = &dto.GlobalFeatureS{Features: make(map[string]dto.Feature)}
	for k, v := range request.Features {
		//datasource.SetFeature(k, v)
		f := dto.Feature{Name: k, Value: v}
		global.Features.Set(f)
	}

	fmt.Println(global.Features)

	flow := request.Flow
	dsl := dslparser.LoadDslFromFile("yaml/" + flow + ".yaml")
	rs := dsl.Parse(global.DslResult)
	c.JSON(http.StatusOK, gin.H{
		"flow":   flow,
		"result": rs,
	})
}
