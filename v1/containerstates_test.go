package conainerstates

import "testing"

func TestAggregate(t *testing.T) {
	table := []struct {
		State1        State
		State2        State
		ExpectedState State
	}{
		{Down, Down, Down},
		{Up, Up, Up},
		{Up, Starting, Starting},
		{Starting, Up, Starting},
		{Up, Stopping, Stopping},
		{Stopping, Up, Stopping},
		{Down, Stopping, Down},
		{Stopping, Down, Down},
		{Failed, Starting, Failed},
		{Starting, Failed, Failed},
		{Failed, Down, Failed},
		{Down, Failed, Failed},
		{Down, Up, Down},
		{Up, Down, Down},
		{Up, Deleted, Deleted},
		{Deleted, Up, Deleted},
		{Deleting, Deleted, Deleting},
		{Deleted, Deleting, Deleting},
		{Deleted, Down, Down},
		{Down, Deleted, Down},
	}

	for index, testEntry := range table {
		agg := Aggregate(testEntry.State1, testEntry.State2)

		if agg != testEntry.ExpectedState {
			t.Errorf("%02d: Aggregate(%v, %v) = %v, expected %v\n",
				index,
				testEntry.State1,
				testEntry.State2,
				agg,
				testEntry.ExpectedState,
			)
		}
	}
}
