package html

import "testing"

func TestGorillaMux(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestGorillaMux"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GorillaMux()
		})
	}
}
