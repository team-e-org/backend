package db

import "testing"

func TestConnectToMySQL(t *testing.T) {
	_, err := ConnectToMySql("test", 0, "test", "test", "test", "test")
	if err == nil {
		t.Error("Error should occur")
	}
}
