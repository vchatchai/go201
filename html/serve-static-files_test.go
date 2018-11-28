package html

import "testing"

func TestStatic(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestStatic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Static()
		})
	}
}
