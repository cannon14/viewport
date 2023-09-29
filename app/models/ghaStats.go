package models

type GhaStats struct {
	ApplicationName string `json:"applicationName"`
	Environment     string `json:"environment"`
	Event           string `json:"event"`
	Count           uint8  `json:"count"`
}
