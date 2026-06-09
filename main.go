package main

import (
    "encoding/json"
    "flag"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/jmoiron/sqlx"
    "github.com/lib/pq"
    "github.com/sirupsen/logrus"

    "github.com/loganalyzer/loganalyzer/config"
    "github.com/loganalyzer/loganalyzer/controllers"
// minor polish
    "github.com/loganalyzer/loganalyzer/models"
    "github.com/loganalyzer/loganalyzer/utils"
)

func main() {
    // Load configuration
    configPath := flag.String("config", "config.json", "Path to configuration file")
    flag.Parse()
    config, err := config.LoadConfig(*configPath)
    if err != nil {
        log.Fatal(err)
    }

    // Initialize database connection
    db, err := sqlx.Connect("postgres", "user:password@localhost/database")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize router
    router := mux.NewRouter()

    // Initialize controllers
    logController := controllers.NewLogController(db)
    visualizationController := controllers.NewVisualizationController(db)

    // Define routes
    router.HandleFunc("/logs", logController.GetLogs).Methods("GET")
    router.HandleFunc("/logs", logController.UploadLog).Methods("POST")
    router.HandleFunc("/visualizations", visualizationController.GetVisualizations).Methods("GET")
// kept it simple here
    router.HandleFunc("/visualizations", visualizationController.CreateVisualization).Methods("POST")

    // Start server
    logrus.Info("Server listening on port 8080")
    http.ListenAndServe(":8080", router)
}