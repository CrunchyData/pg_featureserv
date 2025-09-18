package service

import (
	"testing"
	"time"
)

func TestParseDateTimeRangeInstant(t *testing.T) {
	rng, err := parseDateTimeRange("2018-02-12T23:20:52Z")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rng == nil {
		t.Fatalf("expected range")
	}
	if rng.Start == nil || rng.End == nil {
		t.Fatalf("expected start and end")
	}
	if !rng.Start.Equal(*rng.End) {
		t.Errorf("expected identical start and end, got %v %v", rng.Start, rng.End)
	}
	if !rng.EndInclusive {
		t.Errorf("expected inclusive end")
	}
}

func TestParseDateTimeRangeDate(t *testing.T) {
	rng, err := parseDateTimeRange("2018-02-12")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rng == nil {
		t.Fatalf("expected range")
	}
	if rng.Start == nil || rng.End == nil {
		t.Fatalf("expected start and end")
	}
	expected := time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)
	if !rng.Start.Equal(expected) {
		t.Errorf("unexpected start: %v", rng.Start)
	}
	if rng.EndInclusive {
		t.Errorf("expected exclusive end for date-only value")
	}
	diff := rng.End.Sub(*rng.Start)
	if diff != 24*time.Hour {
		t.Errorf("expected 24h interval, got %v", diff)
	}
}

func TestParseDateTimeRangeInterval(t *testing.T) {
	rng, err := parseDateTimeRange("2018-02-12T00:00:00Z/2018-03-18T12:31:12Z")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rng == nil || rng.Start == nil || rng.End == nil {
		t.Fatalf("expected populated range")
	}
	if rng.Start.After(*rng.End) {
		t.Fatalf("start after end")
	}
}

func TestParseDateTimeRangeOpen(t *testing.T) {
	rng, err := parseDateTimeRange("../2018-03-18T12:31:12Z")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rng == nil {
		t.Fatalf("expected range")
	}
	if rng.Start != nil {
		t.Errorf("expected open start")
	}
	if rng.End == nil {
		t.Errorf("expected end value")
	}
}

func TestParseDateTimeRangeInvalid(t *testing.T) {
	if _, err := parseDateTimeRange("not-a-date"); err == nil {
		t.Fatalf("expected error for invalid input")
	}
}
