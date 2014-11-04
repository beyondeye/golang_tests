package greeter

import (
	"github.com/darkhelmet/env"
	"time"
)

// Greeting returns a pleasant, semi-useful greeting.
func Greeting() string {
	msg := "Hello world, the time is "
	msg += time.Now().String()
	msg += " and your PATH is "
	msg += env.String("PATH")
	return msg
}

