package restful

import "testing"

func TestPost(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestPost"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Post()
		})
	}
}
