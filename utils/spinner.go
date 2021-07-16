package utils

import (
	"fmt"
	"time"
)

func Spinner(delay time.Duration)  {
	for _, r := range `-\|/` {
		fmt.Printf("\r%c >>> ", r)
		time.Sleep(delay)
	}
}
