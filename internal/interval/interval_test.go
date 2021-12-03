package interval_test

import (
	"testing"
	"time"

	"github.com/daniilty/tinkoff-invest-grpc-gateway/internal/interval"
)

func TestIntervalSplit(t *testing.T) {
	const fiftyMins = 50 * time.Minute

	from, err := time.Parse(time.RFC3339, "2021-05-26T10:00:05Z")
	if err != nil {
		t.Fatal(err)
	}

	to := from.Add(3 * time.Hour)

	i := &interval.Interval{
		From: from,
		To:   to,
	}

	expected := []*interval.Interval{
		{
			From: from,
			To:   from.Add(fiftyMins),
		},
		{
			From: from.Add(fiftyMins),
			To:   from.Add(fiftyMins * 2),
		},
		{
			From: from.Add(fiftyMins * 2),
			To:   from.Add(fiftyMins * 3),
		},
		{
			From: from.Add(fiftyMins * 3),
			To:   to,
		},
	}

	intervals := i.Split(50 * time.Minute)
	if len(intervals) != len(expected) {
		t.Errorf("unexpected results len: %d", len(intervals))

		return
	}

	for i, res := range intervals {
		if !res.From.Equal(expected[i].From) {
			t.Errorf("unexpected output from: %s, expected: %s", res.From, expected[i].From)
		}

		if !res.To.Equal(expected[i].To) {
			t.Errorf("unexpected output to: %s, to: %s", res.To, expected[i].To)
		}
	}
}
