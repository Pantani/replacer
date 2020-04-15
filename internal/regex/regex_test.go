package regex

import (
	"testing"
)

func TestVerifyRegexName(t *testing.T) {
	tests := []struct {
		name   string
		verify string
		want   bool
	}{
		{
			"test valid name",
			"bob",
			true,
		}, {
			"test valid name with numbers",
			"bob09",
			true,
		}, {
			"test invalid 1",
			"bob.09",
			false,
		}, {
			"test invalid 2",
			"bob.com",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidName(tt.verify); got != tt.want {
				t.Errorf("IsValidName() = %v, want %v", got, tt.want)
			}
		})
	}
}
