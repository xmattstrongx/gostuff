package main_test

import (
	"fmt"
	"testing"
	"time"
)

// go test -bench .

// benchmarks for different standard time parse/formatting vs manual SprintF formatting. All examples use LDAP time stamp.

func BenchmarkTimeParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = time.Parse("20060102150405.000Z", "05-08-2016")
	}
}

func BenchmarkTimeFormat(b *testing.B) {
	t, _ := time.Parse("20060102150405.000Z", "05-08-2016")
	for i := 0; i < b.N; i++ {
		_ = t.Format("20060102150405.000Z")
	}
}

func BenchmarkTimeParseAndFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t, _ := time.Parse("20060102150405.000Z", "05-08-2016")
		_ = t.Format("20060102150405.000Z")
	}
}

func BenchmarkTimeSprintfManipulation(b *testing.B) {
	t, _ := time.Parse("20060102150405.000Z", "05-08-2016")
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%04d%02d%02d000000Z", t.Year(), t.Month(), t.Day())
	}
}

func BenchmarkTimeParseAndSprintfManipulation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t, _ := time.Parse("20060102150405.000Z", "05-08-2016")
		_ = fmt.Sprintf("%04d%02d%02d000000Z", t.Year(), t.Month(), t.Day())
	}
}
