package searching

type Response struct {
	Cars    []Car
	Filters map[string]interface{}
}
