package restful

import "testing"

func TestHttpRestGet(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHttpRestGet"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpRestGet()
		})
	}
}
