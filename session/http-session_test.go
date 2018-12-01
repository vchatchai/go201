package session

import "testing"

func TestHttpSession(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHttpSession"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpSession()
		})
	}
}
