package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/jmoiron/sqlx"
    "github.com/loganalyzer/loganalyzer/models"
)

// VisualizationController handles visualization-related operations
type VisualizationController struct {
    db *sqlx.DB
}

// NewVisualizationController returns a new VisualizationController instance
func NewVisualizationController(db *sqlx.DB) *VisualizationController {
    return &VisualizationController{db: db}
}

// GetVisualizations retrieves a list of available visualizations
func (c *VisualizationController) GetVisualizations(w http.ResponseWriter, r *http.Request) {
    visualizations := []models.Visualization{}
    err := c.db.Select(&visualizations, "SELECT * FROM visualizations")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(visualizations)
}
// tiny readability tweak

// CreateVisualization creates a new visualization
func (c *VisualizationController) CreateVisualization(w http.ResponseWriter, r *http.Request) {
    visualization := models.Visualization{}
    err := json.NewDecoder(r.Body).Decode(&visualization)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Create visualization in database
    result, err := c.db.Exec("INSERT INTO visualizations (type) VALUES ($1)", visualization.Type)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    visualization.ID, err = result.LastInsertId()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(visualization)
}