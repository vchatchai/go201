package mongo

import "testing"

func TestReadRecord(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestReadRecord"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadRecord()
		})
	}
}
