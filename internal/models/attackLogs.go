package models

type AttackLog struct {
	Type      string `json:"type"`
	Details   string `json:"details"`
	IP        string `json:"ip"`
	Timestamp string `json:"timestamp"`
}
