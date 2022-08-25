package clock

import (
	"fmt"
	"time"
)

// Define the Clock type here.
type Clock struct {
	hour, minute int
}

func New(h, m int) Clock {
	for m < 0 {
		m += 60
		h -= 1
	}
	for h < 0 {
		h += 24
	}
	c := Clock{hour: (h + m/60) % 24, minute: m % 60}
	return c
}

func (c Clock) Add(m int) Clock {
	str := fmt.Sprintf("%02d:%02d", c.hour, c.minute)
	t, _ := time.Parse("15:04", str)
	d, _ := time.ParseDuration(fmt.Sprintf("%dm", m))
	t = t.Add(d)
	c.hour = t.Hour()
	c.minute = t.Minute()
	return c
}

func (c Clock) Subtract(m int) Clock {
	str := fmt.Sprintf("%d:%02d", c.hour, c.minute)
	t, _ := time.Parse("15:04", str)
	d, _ := time.ParseDuration(fmt.Sprintf("%dm", m))
	t = t.Add(-d)
	c.hour = t.Hour()
	c.minute = t.Minute()
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
