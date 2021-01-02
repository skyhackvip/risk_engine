package main

import (
	"github.com/skyhackvip/risk_engine/api"
)

func main() {
	router := api.InitRouter()
	router.Run(":8889")
}
