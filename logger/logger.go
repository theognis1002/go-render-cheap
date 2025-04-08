package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	fileLogger *log.Logger
	file       *os.File
}

// New creates a new logger that writes to both console and file
func New(action string) (*Logger, error) {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("log", 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	// Create log file with timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := filepath.Join("log", fmt.Sprintf("%s_%s.log", action, timestamp))

	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %v", err)
	}

	fileLogger := log.New(file, "", log.LstdFlags)

	return &Logger{
		fileLogger: fileLogger,
		file:       file,
	}, nil
}

// Close closes the log file
func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

// Log writes a message to both console and log file
func (l *Logger) Log(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)

	// Write to console
	fmt.Println(msg)

	// Write to file with timestamp
	l.fileLogger.Println(msg)
}

// Error writes an error message to both console and log file
func (l *Logger) Error(format string, v ...interface{}) {
	msg := fmt.Sprintf("ERROR: "+format, v...)

	// Write to console in red
	fmt.Printf("\033[31m%s\033[0m\n", msg)

	// Write to file with timestamp
	l.fileLogger.Println(msg)
}
