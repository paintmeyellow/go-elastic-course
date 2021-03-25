package searching

type Response struct {
	Hits struct {
		Hits []struct {
			Score  float64 `json:"_score"`
			Source source  `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
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