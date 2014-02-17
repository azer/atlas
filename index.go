package atlas

import "sort"

type Index struct {
	Welcome bool `json:"welcome"`
	EndPoints []string `json:"endpoints"`
}

func NewIndex (urls Map) *Response {
	patterns := make([]string, len(urls))

	i := 0
	for pattern := range urls {
		patterns[i] = pattern
		i++
	}

	sort.Strings(patterns)

	index := &Index{
		true,
		patterns,
	}

	return Manual(200, index)
}
