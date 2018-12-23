package mongo

import "testing"

func TestConnectMongoDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestConnectMongoDB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConnectMongoDB()
		})
	}
}
