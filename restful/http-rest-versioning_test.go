package restful

import "testing"

func TestVersion(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestVersion"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version()
		})
	}
}
