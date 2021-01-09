package log

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/logging"
)

// GoLogger is the default agora logger
type GoLogger struct{}

var cloudLogger *logging.Logger
var localLogger *log.Logger

// New is the next thing
func New(projectID string) *GoLogger {
	cloudLogger = createGCloudLogger(projectID)
	localLogger = log.New(os.Stdout, "[Local]: ", 0)
	return &GoLogger{}
}

func createGCloudLogger(projectID string) *logging.Logger {
	loggingClient, err := logging.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	logName := projectID + ".log"
	return loggingClient.Logger(logName)
}

// Debug is
func (log *GoLogger) Debug(v ...interface{}) {
	localLogger.Println(v...)
	cloudLogger.StandardLogger(logging.Debug).Println(v...)
}

// Error ...
func (log *GoLogger) Error(e error) {
	localLogger.Println(e)
	cloudLogger.StandardLogger(logging.Error).Println(e)
}
