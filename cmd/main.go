package main

import (
	"fmt"
	"os"

	"github.com/artepepep/Get-stats-GO/conf"
	"github.com/artepepep/Get-stats-GO/elasticsearch"
)

var (
	accountIds = []int{9730}
)

func main() {
	conf.LoadEnv()
	es := elasticsearch.NewClient(os.Getenv("ELASTIC_SEARCH_URL"))
	q := elasticsearch.Query{}

	queryBody := q.BuildAggregation(accountIds, "2025-11-10", "2025-11-16")
	resp, err := es.Search("statistic9730", queryBody)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(resp))
}
