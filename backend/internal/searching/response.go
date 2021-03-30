package searching

import "github.com/paintmeyellow/go-elastic-course/internal/elastic"

type Response struct {
	Cars    []Car
	Filters Filters
}

type Filters elastic.Filters
