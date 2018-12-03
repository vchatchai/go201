package httpsession

import "testing"

func TestRedis(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestRedis"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Redis()
		})
	}
}
