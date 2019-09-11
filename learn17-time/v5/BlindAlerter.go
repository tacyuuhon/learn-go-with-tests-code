package poker

import (
	"fmt"
	"os"
	"time"
)

// BlindAlerter interface
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// BlindAlerterFunc func def
type BlindAlerterFunc func(duration time.Duration, amount int)

// ScheduleAlertAt func
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

// StdOutAlerter func
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
