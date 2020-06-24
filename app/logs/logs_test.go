package logs

import (
	"testing"
)

func TestLogger(t *testing.T) {
	if logger == nil {
		t.Fatalf("logger instance does not exist")
	}
}

func TestInfoLogger(t *testing.T) {
	Info("test")
}

func TestWarnLogger(t *testing.T) {
	Warn("test")
}

func TestErrorLogger(t *testing.T) {
	Error("test")
}

func TestDebugLogger(t *testing.T) {
	Debug("test")
}
