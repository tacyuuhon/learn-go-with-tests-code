package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const sleep = "sleep"
const write = "write"

// Sleeper intefrace
type Sleeper interface {
	Sleep()
}

// SpySleeper struct
type SpySleeper struct {
	Calls int
}

// Sleep func
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// ConfigurableSleeper struct
type ConfigurableSleeper struct {
	duration time.Duration
}

// Sleep func
func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

// CountdownOperationsSpy struct
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleep func
func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

// Countdown func
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
