package tests

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/jmoiron/sqlx"
    "github.com/loganalyzer/loganalyzer/controllers"
    "github.com/loganalyzer/loganalyzer/models"
)

func TestGetLogs(t *testing.T) {
    // Create test database
    db, err := sqlx.Connect("postgres", "user:password@localhost/database")
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()

    // Create log entries
    logs := []models.Log{
        {Message: "Log 1"},
        {Message: "Log 2"},
    }
    for _, log := range logs {
        _, err := db.Exec("INSERT INTO logs (message) VALUES ($1)", log.Message)
        if err != nil {
            t.Fatal(err)
        }
    }

    // Create log controller
    logController := controllers.NewLogController(db)

    // Create test request
    req, err := http.NewRequest("GET", "/logs", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create test response recorder
    w := httptest.NewRecorder()

    // Call log controller
    logController.GetLogs(w, req)

    // Check response status code
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    // Check response body
    var response []models.Log
    err = json.Unmarshal(w.Body.Bytes(), &response)
    if err != nil {
        t.Fatal(err)
    }

    if len(response) != len(logs) {
        t.Errorf("Expected %d log entries, got %d", len(logs), len(response))
    }
}