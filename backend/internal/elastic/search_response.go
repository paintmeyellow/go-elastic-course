package elastic

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			Score  float64 `json:"_score"`
			Source struct {
				ID     int    `json:"id"`
				Make   string `json:"make"`
				Model  string `json:"model"`
				Params struct {
					Color string  `json:"color"`
					Year  uint16  `json:"year"`
					Price float64 `json:"price"`
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

func (req SearchResponse) Filters() map[string]interface{} {
	var (
		filters  = map[string]interface{}{}
		minPrice = req.Aggs.MinPrice
		maxPrice = req.Aggs.MaxPrice
		color    = req.Aggs.Color
		year     = req.Aggs.Year
	)

	if minPrice.Value() != maxPrice.Value() {
		filters["Price"] = rangeFilter{
			Min: minPrice.Value(),
			Max: maxPrice.Value(),
		}.build()
	}

	if !color.isEmpty() {
		filters["Color"] = checkboxFilter{
			Variants: color.variants(),
		}.build()
	}

	if !year.isEmpty() {
		filters["Year"] = checkboxFilter{
			Variants: year.variants(),
		}.build()
	}

	return filters
}

/****************
**** Filters ****
*****************/

type checkboxFilter struct {
	Variants []checkboxFilterVariant
}

type checkboxFilterVariant struct {
	Value interface{}
	Count uint
}

func (f checkboxFilter) build() map[string]interface{} {
	vv := make([]map[string]interface{}, 0)

	for _, v := range f.Variants {
		vv = append(vv, map[string]interface{}{
			"value": v.Value,
			"count": v.Count,
		})
	}

	return map[string]interface{}{
		"type":     "checkbox",
		"variants": vv,
	}
}

type rangeFilter struct {
	Min float64
	Max float64
}

func (f rangeFilter) build() map[string]interface{} {
	return map[string]interface{}{
		"type": "range",
		"min":  f.Min,
		"max":  f.Max,
	}
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

func (agg termsAggregation) variants() []checkboxFilterVariant {
	var vv []checkboxFilterVariant
	for _, b := range agg.Vars.Buckets {
		vv = append(vv, checkboxFilterVariant{
			Value: b.Key,
			Count: b.DocCount,
		})
	}
	return vv
}

func (agg termsAggregation) isEmpty() bool {
	return len(agg.Vars.Buckets) == 0
}
