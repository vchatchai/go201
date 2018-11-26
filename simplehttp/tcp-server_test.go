package simplehttp

import "testing"

func TestTcpServer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestTcpServer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TcpServer()
		})
	}
}
