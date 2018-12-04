package main

import (
	"testing"
)

func TestGuardMostAsleepWithMinute(t *testing.T) {
	for n, tc := range []struct {
		in     []string
		guard  int
		minute int
	}{
		{
			[]string{
				"[1518-11-01 00:00] Guard #10 begins shift",
				"[1518-11-01 00:05] falls asleep",
				"[1518-11-01 00:25] wakes up",
				"[1518-11-01 00:30] falls asleep",
				"[1518-11-01 00:55] wakes up",
				"[1518-11-01 23:58] Guard #99 begins shift",
				"[1518-11-02 00:40] falls asleep",
				"[1518-11-02 00:50] wakes up",
				"[1518-11-03 00:05] Guard #10 begins shift",
				"[1518-11-03 00:24] falls asleep",
				"[1518-11-03 00:29] wakes up",
				"[1518-11-04 00:02] Guard #99 begins shift",
				"[1518-11-04 00:36] falls asleep",
				"[1518-11-04 00:46] wakes up",
				"[1518-11-05 00:03] Guard #99 begins shift",
				"[1518-11-05 00:45] falls asleep",
				"[1518-11-05 00:55] wakes up",
			},
			10, 24,
		},
	} {
		guard, minute := guardMostAsleepWithMinute(tc.in)

		if got, want := guard, tc.guard; got != want {
			t.Errorf("[%d] guard = %d, want %d", n, got, want)
		}

		if got, want := minute, tc.minute; got != want {
			t.Errorf("[%d] minute = %d, want %d", n, got, want)
		}
	}
}

func TestGuardMostAsleepWithMinuteAlternate(t *testing.T) {
	for n, tc := range []struct {
		in     []string
		guard  int
		minute int
	}{
		{
			[]string{
				"[1518-11-01 00:00] Guard #10 begins shift",
				"[1518-11-01 00:05] falls asleep",
				"[1518-11-01 00:25] wakes up",
				"[1518-11-01 00:30] falls asleep",
				"[1518-11-01 00:55] wakes up",
				"[1518-11-01 23:58] Guard #99 begins shift",
				"[1518-11-02 00:40] falls asleep",
				"[1518-11-02 00:50] wakes up",
				"[1518-11-03 00:05] Guard #10 begins shift",
				"[1518-11-03 00:24] falls asleep",
				"[1518-11-03 00:29] wakes up",
				"[1518-11-04 00:02] Guard #99 begins shift",
				"[1518-11-04 00:36] falls asleep",
				"[1518-11-04 00:46] wakes up",
				"[1518-11-05 00:03] Guard #99 begins shift",
				"[1518-11-05 00:45] falls asleep",
				"[1518-11-05 00:55] wakes up",
			},
			99, 45,
		},
	} {
		guard, minute := guardMostAsleepWithMinuteAlternate(tc.in)

		if got, want := guard, tc.guard; got != want {
			t.Errorf("[%d] guard = %d, want %d", n, got, want)
		}

		if got, want := minute, tc.minute; got != want {
			t.Errorf("[%d] minute = %d, want %d", n, got, want)
		}
	}
}
