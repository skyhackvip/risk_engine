package handler

import (
	"github.com/gin-gonic/gin"
	//	"github.com/skyhackvip/global"
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/internal"
	"net/http"
	"strconv"
)

func DslRunHandler(c *gin.Context) {
	f1 := c.PostForm("f1")
	f2 := c.PostForm("f2")
	f_1, _ := strconv.ParseInt(f1, 0, 64)
	f_2, _ := strconv.ParseBool(f2)
	internal.SetFeature("feature_1", f_1)
	internal.SetFeature("feature_2", f_2)

	dsl := dslparser.LoadDslFromFile("decisiontree.yaml")
	//var dslResult global.DslResult
	//	dsl.SetResult(&dslResult)

	//rs := dsl.ParseDecisionTree(dsl.Decisiontrees[0])
	rs := dsl.Parse()

	c.JSON(http.StatusOK, gin.H{
		"rs": rs,
		"f1": f_1,
		"f2": f_2,
	})
}
