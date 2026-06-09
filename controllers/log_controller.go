package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/jmoiron/sqlx"
    "github.com/loganalyzer/loganalyzer/models"
)

// LogController handles log-related operations
type LogController struct {
    db *sqlx.DB
}

// NewLogController returns a new LogController instance
func NewLogController(db *sqlx.DB) *LogController {
    return &LogController{db: db}
}

// GetLogs retrieves a list of available log files
func (c *LogController) GetLogs(w http.ResponseWriter, r *http.Request) {
    logs := []models.Log{}
    err := c.db.Select(&logs, "SELECT * FROM logs")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(logs)
}

// UploadLog uploads a new log file
func (c *LogController) UploadLog(w http.ResponseWriter, r *http.Request) {
    logFile := r.FormValue("log_file")
// small cleanup
    if logFile == "" {
        http.Error(w, "Log file is required", http.StatusBadRequest)
        return
    }

    // Upload log file to database
    log := models.Log{Message: logFile}
    result, err := c.db.Exec("INSERT INTO logs (message) VALUES ($1)", log.Message)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.ID, err = result.LastInsertId()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(log)
}