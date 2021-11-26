package server

import (
	"fmt"
	"time"
)

func validateIntervals(from time.Time, to time.Time, interval int64) error {
	if interval == 0 {
		return fmt.Errorf("interval cannot be 0")
	}

	if from.After(to) {
		return fmt.Errorf("from cannot be after to")
	}

	if to.Unix()-from.Unix() < interval {
		return fmt.Errorf("interval is too big")
	}

	return nil
}
