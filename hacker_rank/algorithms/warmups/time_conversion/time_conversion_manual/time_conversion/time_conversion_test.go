package time_conversion

import "testing"

// go test -bench .

// benchmarks for different standard time parse/formatting vs manual SprintF formatting. All examples use LDAP time stamp.

func BenchmarkConvertTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ConvertTime("07:05:45PM")
	}
}

func TestTimeParseHappyPath(t *testing.T) {
	if have, want := ConvertTime("07:05:45PM"), "19:05:45"; have != want {
		t.Errorf("expected %s but have %s", want, have)
	}
}
