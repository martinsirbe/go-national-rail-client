package models

import "time"

type DepartureInfo struct {
	OriginCode       string             `json:"origin_code"`
	Origin           string             `json:"origin"`
	DestinationCode  string             `json:"destination_code"`
	Destination      string             `json:"destination"`
	WarningMessages  []string           `json:"warning_messages,omitempty"`
	DepartureDetails []DepartureDetails `json:"departure_details"`
}

type DepartureDetails struct {
	OriginLocation      string     `json:"origin_location"`
	DestinationLocation string     `json:"destination_location"`
	Via                 *string    `json:"via,omitempty"`
	Time                *time.Time `json:"departure_time,omitempty"`
	Status              string     `json:"status"`
	Operator            string     `json:"operator"`
	Platform            string     `json:"platform"`
}
