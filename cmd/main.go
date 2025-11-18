package main

import (
	"fmt"
	"os"

	"github.com/artepepep/Get-stats-GO/conf"
	"github.com/artepepep/Get-stats-GO/utils"
)

var (
	accountIds = []int{9730}
)

func main() {
	conf.LoadEnv()
	es := utils.NewES(os.Getenv("ELASTIC_SEARCH_URL"))
	q := utils.ESQuery{}

	queryBody := q.BuildAggregation(accountIds, "2025-11-10", "2025-11-16")
	resp, err := es.Search(queryBody)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(resp))
}
