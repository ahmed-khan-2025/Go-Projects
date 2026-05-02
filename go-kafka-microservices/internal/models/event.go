package models

type Event struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Payload   string `json:"payload"`
	Timestamp string `json:"timestamp"`
}