package models

import (
    "time"
)

// Log represents a log entry
type Log struct {
    ID        int64     `json:"id"`
    Message   string    `json:"message"`
    Timestamp time.Time `json:"timestamp"`
}