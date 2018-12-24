package mongo

import "testing"

func TestOneRecord(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestOneRecord"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OneRecord()
		})
	}
}
