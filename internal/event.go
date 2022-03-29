package awards

import (
	"time"
)

// Represents an interaction (or the lack of) by a user with a test pishing email.
// Opened: Opened the test pishing email.
// Clicked: Clicked on the phishing link in the test pishing email.
// Reported: Forwarded the test pishing email to the malware-scanner.
// Ignored: The test pishing email was not interacted with by the user for a period of time.
type PhishingAction int

const (
	Opened PhishingAction = iota
	Clicked
	Reported
	Ignored
)

func (p PhishingAction) ToAwardType() AwardType {
	switch p {
	case Opened:
		return OpenAward
	case Reported:
		return ReportAward
	case Ignored:
		return IgnoreAward
	default:
		return -1
	}
}

func (a PhishingAction) String() string {
	switch a {
	case Opened:
		return "opened"
	case Reported:
		return "reported"
	case Ignored:
		return "ignored"
	case Clicked:
		return "clicked"
	default:
		return ""
	}
}

func ToPhishingAction(action string) PhishingAction {
	switch action {
	case "opened":
		return Opened
	case "reported":
		return Reported
	case "ignored":
		return Ignored
	case "clicked":
		return Clicked
	default:
		return -1
	}
}

// A user interaction with a phishing mail.
type UserPhishingEvent struct {
	ID          int64
	UserID      UserId
	Action      PhishingAction
	CreatedAt   time.Time
	EmailRef    string
	ProcessedAt *time.Time
}
