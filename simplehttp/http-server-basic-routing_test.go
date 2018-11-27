package simplehttp

import "testing"

func TestRouting(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestRouting"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Routing()
		})
	}
}
