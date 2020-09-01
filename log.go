package log

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/logging"
	"github.com/joho/godotenv"
)

// GoLogger is the default agora logger
type GoLogger interface {
	Debug(json []byte)
	Error(e error)
}

// Logger ...
var Logger *GoLogger

var cloudLogger *logging.Logger
var localLogger *log.Logger

func init() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	gCloudProjectID := os.Getenv("GOOGLE_CLOUD_PROJECT_ID")
	cloudLogger = createGCloudLogger(gCloudProjectID)
	localLogger = log.New(os.Stdout, "[Local]: ", 0)
}

func createGCloudLogger(projectID string) *logging.Logger {
	loggingClient, err := logging.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	logName := "jaspergif-log"
	return loggingClient.Logger(logName)
}

// Debug is
func Debug(v ...interface{}) {
	localLogger.Println(v...)
	cloudLogger.StandardLogger(logging.Debug).Println(v...)
}

// Error ...
func Error(e error) {
	localLogger.Println(e)
	cloudLogger.StandardLogger(logging.Error).Println(e)
}
