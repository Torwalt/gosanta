package events

import "time"

type PhishingEvent struct {
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Action    string    `json:"action"`
	EmailRef  string    `json:"email_ref"`
	EventId   string
}
