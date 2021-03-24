package indexing

type Car struct {
	ID     int    `json:"id"`
	Make   string `json:"make"`
	Model  string `json:"model"`
	Params struct {
		Color string  `json:"color"`
		Year  uint16  `json:"year"`
		Price float64 `json:"price"`
	} `json:"params"`
}
