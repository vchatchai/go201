package simplehttp

import "testing"

func TestHttpServerBasicAuthentication(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHttpServerBasicAuthentication"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpServerBasicAuthentication()
		})
	}
}
