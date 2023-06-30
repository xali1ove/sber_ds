package database

import (
	"database/sql"
	"fmt"
	"sber_ds/models"
	"time"
)

func GetReport(db *sql.DB, reportID int) (*models.Report, error) {
	query := "SELECT * FROM reports WHERE id = $1"
	row := db.QueryRow(query, reportID)

	report := &models.Report{}
	err := row.Scan(&report.ID, &report.ReportInfo, &report.CreationTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("report not found")
		}
		return nil, err
	}

	return report, nil
}

func InsertReport(db *sql.DB, reportInfo string) error {
	query := "INSERT INTO reports (report_info, creation_time) VALUES ($1, $2)"
	_, err := db.Exec(query, reportInfo, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func GetObservationTime(db *sql.DB, modelID int) (int, error) {
	query := "SELECT MAX(EXTRACT(DAY FROM NOW() - creation_time)) FROM reports WHERE model_id = $1"
	row := db.QueryRow(query, modelID)

	var observationPeriod int
	err := row.Scan(&observationPeriod)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no reports found for the model")
		}
		return 0, err
	}

	return observationPeriod, nil
}
