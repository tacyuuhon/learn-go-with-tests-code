package poker

import "time"

// BlindAlerter interface
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}
