package utils

import "time"

type AlarmData struct {
	Time    time.Time `json:"time"`
	Name    string    `json:"name"`
	Message string    `json:"message"`
}
type AlarmInfo struct {
	Name    string
	Time    string
	Message string
	Stop    chan string
}
