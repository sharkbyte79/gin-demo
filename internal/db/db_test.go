package db

import (
	"context"
	"testing"
)

// Test establishing a successful connection to the database
func TestConnect(t *testing.T) {
	_, err := Connect()

	if err != nil {
		t.Errorf("got %v", err)
	}
}

// Test if pinging the database returns an error
func TestConnectPing(t *testing.T) {
	conn, err := Connect()

	if err != nil || conn.Ping(context.Background()) != nil {
		t.Errorf("got %v", err)
	}
	defer conn.Close(context.TODO())
}
