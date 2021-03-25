package searching

type Car struct {
	ID     int    `json:"id"`
	Make   string `json:"make"`
	Model  string `json:"model"`
	Params Params `json:"params"`
}

type Params struct {
	Color string  `json:"color"`
	Year  uint16  `json:"year"`
	Price float64 `json:"price"`
}
