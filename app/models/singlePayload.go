package models

type SinglePayload struct {
	Draw            int        `json:"draw" default:"10"`
	RecordsTotal    int64      `json:"recordsTotal" default:"10"`
	RecordsFiltered int64      `json:"recordsFiltered" default:"10"`
	Data            AwsAccount `json:"data"`
	Error           string     `json:"error"`
}
