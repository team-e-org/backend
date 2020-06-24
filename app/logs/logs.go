package logs

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
}

func Info(msg string, a ...interface{}) {
	logger.Printf("[INFO] %s\n", fmt.Sprintf(msg, a...))
}

func Warn(msg string, a ...interface{}) {
	logger.Printf("[WARN] %s\n", fmt.Sprintf(msg, a...))
}

func Error(msg string, a ...interface{}) {
	logger.Printf("[ERROR] %s\n", fmt.Sprintf(msg, a...))
}

func Debug(msg string, a ...interface{}) {
	logger.Printf("[DEBUG] %s\n", fmt.Sprintf(msg, a...))
}
