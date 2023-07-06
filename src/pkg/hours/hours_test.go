package hours

import (
	"testing"
	"time"
)

func TestTimeDifference(t *testing.T) {
	tests := []struct {
		t1       string
		t2       string
		expected time.Duration
	}{
		{"12:00", "13:00", time.Hour},
		{"23:00", "01:00", 2 * time.Hour},
		{"01:00", "23:00", 22 * time.Hour},
		{"12:00", "12:00", 0},
	}

	for _, test := range tests {
		diff, err := TimeDifference(test.t1, test.t2)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if diff != test.expected {
			t.Errorf("TimeDifference(%s, %s) = %v, expected %v", test.t1, test.t2, diff, test.expected)
		}
	}
}

func TestTimeDifference_Error(t *testing.T) {
	_, err := TimeDifference("12:00", "24:00")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

