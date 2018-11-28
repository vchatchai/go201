package html

import "testing"

func TestFirstTemplate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestFirstTemplate"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FirstTemplate()
		})
	}
}
