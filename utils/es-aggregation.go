package utils

type ESQuery struct {
	Size  int         			`json:"size"`
	Query map[string]any  		`json:"query"`
	Aggs  map[string]any        `json:"aggs"`
}

func NewES(url string) *ESCmd {
	return &ESCmd{
		Url: url,
	}
}

func (aggr *ESQuery) BuildAggregation(accountIds []int, from string, to string) *ESQuery {
	aggr.Size = 0

	aggr.Query = map[string]any{
		"bool": map[string]any{
			"must": []any{
				map[string]any{
					"terms": map[string]any{
						"accountid": accountIds,
					},
				},
				map[string]any{
					"range": map[string]any{
						"date_time": map[string]any{
							"gte": from,
							"lte": to,
						},
					},
				},
			},
		},
	}

	aggr.Aggs = map[string]any{
		"by_site": map[string]any{
			"terms": map[string]any{
				"field": "sitename.keyword",
				"size":  200,
			},
			"aggs": map[string]any{
				"total_revenue": map[string]any{
					"sum": map[string]string{
						"field": "pubCPM",
					},
				},
				"total_impressions": map[string]any{
					"sum": map[string]string{
						"field": "imp_h",
					},
				},
			},
		},
	}

	return aggr
}