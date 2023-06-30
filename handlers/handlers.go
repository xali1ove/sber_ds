package handlers

import (
	"net/http"
	"sber_ds/config"
	"sber_ds/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReport(c *gin.Context) {
	config, _ := config.LoadConfig()
	db, err := database.NewDBConnection(config)
	if err != nil {
		return
	}

	reportID, _ := strconv.Atoi(c.Query("report_id"))

	report, err := database.GetReport(db, reportID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error_msg": err.Error(), "report_info": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error_msg": "", "report_info": report.ReportInfo})
}

func SetReport(c *gin.Context) {
	config, _ := config.LoadConfig()
	db, err := database.NewDBConnection(config)
	if err != nil {
		return
	}
	var request struct {
		ReportInfo string `json:"report_info"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_msg": err.Error()})
		return
	}

	if err := database.InsertReport(db, request.ReportInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error_msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error_msg": ""})
}

func GetObservationTime(c *gin.Context) {
	config, _ := config.LoadConfig()
	db, err := database.NewDBConnection(config)
	if err != nil {
		return
	}
	modelID, _ := strconv.Atoi(c.Query("model_id"))

	observationPeriod, err := database.GetObservationTime(db, modelID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error_msg": err.Error(), "max_observation_period": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error_msg": "", "max_observation_period": strconv.Itoa(observationPeriod)})
}
