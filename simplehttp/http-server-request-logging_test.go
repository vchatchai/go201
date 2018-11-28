package simplehttp

import "testing"

func TestLog(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestLog"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Log()
		})
	}
}
