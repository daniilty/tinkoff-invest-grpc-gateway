package interval

import "time"

type Interval struct {
	From time.Time
	To   time.Time
}

func (i *Interval) Split(duration time.Duration) []*Interval {
	if isFitsDuration(i.From, i.To, duration) {
		return []*Interval{i}
	}

	splitted := []*Interval{}

	counter := i.From

	for !isFitsDuration(counter, i.To, duration) {
		splitted = append(splitted, &Interval{
			From: counter,
			To:   counter.Add(duration),
		})

		counter = counter.Add(duration)
	}

	if counter.Before(i.To) {
		splitted = append(splitted, &Interval{
			From: counter,
			To:   i.To,
		})
	}

	return splitted
}

func isFitsDuration(from time.Time, to time.Time, duration time.Duration) bool {
	return to.Sub(from) <= duration
}
