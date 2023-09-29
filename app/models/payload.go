package models

type Payload[P any] struct {
	Draw            int    `json:"draw" default:"10"`
	RecordsTotal    int64  `json:"recordsTotal" default:"10"`
	RecordsFiltered int64  `json:"recordsFiltered" default:"10"`
	Data            []P    `json:"data"`
	Error           string `json:"error"`
}
