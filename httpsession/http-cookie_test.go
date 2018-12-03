package httpsession

import "testing"

func TestCookie(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestCookie"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Cookie()
		})
	}
}
