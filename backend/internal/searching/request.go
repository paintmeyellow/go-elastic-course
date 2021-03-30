package searching

type Request struct {
	Index         string
	Query         string
	ActiveFilters ActiveFilters
}

type ActiveFilters struct {
	Checkbox map[string][]string
	Range    map[string]struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	}
}
