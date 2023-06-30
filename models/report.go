package models

import "time"

type Report struct {
	ID           int       `json:"id"`
	ReportInfo   string    `json:"report_info"`
	CreationTime time.Time `json:"creation_time"`
}
