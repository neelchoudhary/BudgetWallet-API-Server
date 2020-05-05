package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LogIfFatalAndExit if error, logs and exits
func LogIfFatalAndExit(err error, msgs ...string) {
	if err != nil {
		var message string
		for _, msg := range msgs {
			message += msg + " "
		}
		log.Fatal(message + ": " + err.Error())
	}
}

// InternalServerError grpc error to client
var InternalServerError = status.Errorf(codes.Internal, "Internal Server Error")

// StartTxErrorMsg error text
var StartTxErrorMsg = "Failed to start new db transaction"

// CommitTxErrorMsg error text
var CommitTxErrorMsg = "Failed to commit db transaction"

// FuncCallErrorMsg error text
var FuncCallErrorMsg = func(callerName string, methodName string) string {
	return callerName + " call to " + methodName + " failed"
}

// InitializeLogs setup logging agent
func InitializeLogs() {
	log.SetFormatter(&log.TextFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.TraceLevel)
}

// Log fields: time, service, repo (if applicable), model (if applicable), method name.
// Trace, Debug, Info, Warn, Error, FATAL, PANIC
// Log msg message: GetAllItems error, failed to
