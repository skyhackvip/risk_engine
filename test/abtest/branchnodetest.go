package main

import (
	"fmt"
	"github.com/skyhackvip/risk_engine/dslparser"
)

func main() {
	dsl := dslparser.LoadDslFromFile("branchnode.yaml")
	fmt.Println("start")
	fmt.Println(dsl)
	//fmt.Println(dsl.ParseBranchnode(dsl.Branchnodes[0]))
	dsl.Parse()
}
