package models

import "time"

type DepartureInfo struct {
	OriginLocationCode      string             `json:"origin_code"`
	DestinationLocationCode string             `json:"destination_code"`
	WarningMessages         []string           `json:"warning_messages,omitempty"`
	DepartureDetails        []DepartureDetails `json:"departure_details"`
}

type DepartureDetails struct {
	OriginLocation string     `json:"origin_location"`
	Destination    string     `json:"destination_location"`
	Via            *string    `json:"via,omitempty"`
	Time           *time.Time `json:"departure_time,omitempty"`
	Status         string     `json:"status"`
	Operator       string     `json:"operator"`
	Platform       string     `json:"platform"`
}
