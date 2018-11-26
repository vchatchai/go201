package simplehttp

import "testing"

func TestHttpServer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHttpServer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpServer()
		})
	}
}
