package data

import (
	"testing"
	"time"
)

func TestSQLDateTimeFilterInstant(t *testing.T) {
	start := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)
	Column := "observed_at"
	rng := &TimeRange{
		Start:          &start,
		End:            &start,
		StartInclusive: true,
		EndInclusive:   true,
	}
	sql, args := sqlDateTimeFilter(Column, Column, rng, 1)
	expected := "(\"observed_at\" IS NULL OR (\"observed_at\" >= $1 AND \"observed_at\" <= $2))"
	if sql != expected {
		t.Fatalf("unexpected sql: \n%s\n%s", sql, expected)
	}
	if len(args) != 2 {
		t.Fatalf("expected 2 args, got %d", len(args))
	}
	if !args[0].(time.Time).Equal(start) {
		t.Errorf("unexpected start arg: %v", args[0])
	}
	if !args[1].(time.Time).Equal(start) {
		t.Errorf("unexpected end arg: %v", args[1])
	}
}

func TestSQLDateTimeFilterExclusiveEnd(t *testing.T) {
	start := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)
	end := start.Add(24 * time.Hour)
	Column := "observed_at"
	rng := &TimeRange{
		Start:          &start,
		End:            &end,
		StartInclusive: true,
		EndInclusive:   false,
	}
	sql, args := sqlDateTimeFilter(Column, Column, rng, 3)
	expected := "(\"observed_at\" IS NULL OR (\"observed_at\" >= $3 AND \"observed_at\" < $4))"
	if sql != expected {
		t.Fatalf("unexpected sql: %s", sql)
	}
	if len(args) != 2 {
		t.Fatalf("expected 2 args, got %d", len(args))
	}
	if !args[0].(time.Time).Equal(start) {
		t.Errorf("unexpected start arg: %v", args[0])
	}
	if !args[1].(time.Time).Equal(end) {
		t.Errorf("unexpected end arg: %v", args[1])
	}
}

func TestSQLDateTimeFilterNilRange(t *testing.T) {
	sql, args := sqlDateTimeFilter("", "", nil, 1)
	if sql != "" {
		t.Fatalf("expected empty sql")
	}
	if args != nil {
		t.Fatalf("expected nil args")
	}
}

func TestSQLDateTimeFilterIntervalColumns(t *testing.T) {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC)
	StartColumn := "start_time"
	EndColumn := "end_time"
	rng := &TimeRange{
		Start:          &start,
		End:            &end,
		StartInclusive: true,
		EndInclusive:   true,
	}
	sql, args := sqlDateTimeFilter(StartColumn, EndColumn, rng, 2)
	expected := "(\"end_time\" IS NULL OR \"end_time\" >= $2) AND (\"start_time\" IS NULL OR \"start_time\" <= $3)"
	if sql != expected {
		t.Fatalf("unexpected sql: %s", sql)
	}
	if len(args) != 2 {
		t.Fatalf("expected 2 args, got %d", len(args))
	}
	if !args[0].(time.Time).Equal(start) {
		t.Errorf("unexpected start arg: %v", args[0])
	}
	if !args[1].(time.Time).Equal(end) {
		t.Errorf("unexpected end arg: %v", args[1])
	}
}
