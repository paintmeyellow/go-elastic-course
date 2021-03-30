package elastic

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			Score  float64 `json:"_score"`
			Source struct {
				ID     int     `json:"id"`
				Make   string  `json:"make"`
				Model  string  `json:"model"`
				Price  float64 `json:"price"`
				Params struct {
					Color string `json:"color"`
					Year  uint16 `json:"year"`
				} `json:"params"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
	Aggs struct {
		MinPrice edgeAggregation  `json:"min_price"`
		MaxPrice edgeAggregation  `json:"max_price"`
		Color    termsAggregation `json:"color"`
		Year     termsAggregation `json:"year"`
	} `json:"aggregations"`
}

/****************
**** Filters ****
*****************/

type Filters struct {
	Checkbox []checkboxFilter `json:"checkbox"`
	Range    []rangeFilter    `json:"range"`
}

type checkboxFilter struct {
	Name  string               `json:"name"`
	Items []checkboxFilterItem `json:"items"`
}

type checkboxFilterItem struct {
	Value string `json:"value"`
	Count uint   `json:"count"`
}

type rangeFilter struct {
	Name string  `json:"name"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
}

func (req SearchResponse) Filters() Filters {
	var (
		filters  Filters
		minPrice = req.Aggs.MinPrice
		maxPrice = req.Aggs.MaxPrice
		color    = req.Aggs.Color
		year     = req.Aggs.Year
	)

	if minPrice.Value() != maxPrice.Value() {
		filters.Range = append(filters.Range, rangeFilter{
			Name: "Price",
			Min:  minPrice.Value(),
			Max:  maxPrice.Value(),
		})
	}

	if !color.isEmpty() {
		filters.Checkbox = append(filters.Checkbox, checkboxFilter{
			Name:  "Color",
			Items: color.checkboxFilterItems(),
		})
	}

	if !year.isEmpty() {
		filters.Checkbox = append(filters.Checkbox, checkboxFilter{
			Name:  "Year",
			Items: year.checkboxFilterItems(),
		})
	}

	return filters
}

/*********************
**** Aggregations ****
**********************/

type edgeAggregation struct {
	Val struct {
		Value float64 `json:"value"`
	} `json:"val"`
}

func (p edgeAggregation) Value() float64 {
	return p.Val.Value
}

type termsAggregation struct {
	Vars struct {
		Buckets []struct {
			Key      string `json:"key"`
			DocCount uint   `json:"doc_count"`
		} `json:"buckets"`
	} `json:"vars"`
}

func (agg termsAggregation) checkboxFilterItems() []checkboxFilterItem {
	var items []checkboxFilterItem
	for _, b := range agg.Vars.Buckets {
		items = append(items, checkboxFilterItem{
			Value: b.Key,
			Count: b.DocCount,
		})
	}
	return items
}

func (agg termsAggregation) isEmpty() bool {
	return len(agg.Vars.Buckets) == 0
}
