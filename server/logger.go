package server

import (
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// FTPLogger is a logger implementation for use in `go-ftp`.
type FTPLogger struct{}

// Print logs the given message and session id.
func (logger *FTPLogger) Print(sessionID string, message interface{}) {
	logrus.WithFields(logrus.Fields{"time": time.Now(), "session": sessionID, "message": message}).Debug("FTP:", message)
}

// PrintCommand logs the given command and its parameters as well as the session id.
func (logger *FTPLogger) PrintCommand(sessionID string, command string, params string) {
	logrus.WithFields(logrus.Fields{"time": time.Now(), "session": sessionID, "command": command, "parameters": params}).Debugf("FTP: %s(%s)", command, params)
}

// PrintResponse logs the response code and message as well as the session id.
func (logger *FTPLogger) PrintResponse(sessionID string, code int, message string) {
	logrus.WithFields(logrus.Fields{"time": time.Now(), "session": sessionID, "code": code, "response": message}).Debugf("Response with %q and code %d", message, code)

}

// Printf logs an evaluated format string with arguments and the given session id.
func (logger *FTPLogger) Printf(sessionID string, format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	logrus.WithFields(logrus.Fields{"time": time.Now(), "session": sessionID}).Debugf(format, v...)
}
