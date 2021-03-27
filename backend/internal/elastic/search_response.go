package elastic

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			Score  float64 `json:"_score"`
			Source source  `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
	Aggs struct {
		MinPrice price `json:"min_price"`
		MaxPrice price `json:"max_price"`
	} `json:"aggregations"`
}

func (req SearchResponse) Filters() []map[string]interface{} {
	return []map[string]interface{}{
		rangeFilter{
			Name: "Price",
			Min:  req.Aggs.MinPrice.Value(),
			Max:  req.Aggs.MaxPrice.Value(),
		}.build(),
	}
}

type rangeFilter struct {
	Name string
	Min  float64
	Max  float64
}

func (f rangeFilter) build() map[string]interface{} {
	return map[string]interface{}{
		f.Name: map[string]interface{}{
			"type": "range",
			"min":  f.Min,
			"max":  f.Max,
		},
	}
}

type source struct {
	ID     int    `json:"id"`
	Make   string `json:"make"`
	Model  string `json:"model"`
	Params struct {
		Color string  `json:"color"`
		Year  uint16  `json:"year"`
		Price float64 `json:"price"`
	} `json:"params"`
}

type price struct {
	Price struct {
		Value float64 `json:"value"`
	} `json:"price"`
}

func (p price) Value() float64 {
	return p.Price.Value
}
