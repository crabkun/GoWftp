package server

import (
	"fmt"
	"log"
)

// Use an instance of this to log in a standard format
type Logger struct {
	sessionID string
	show bool
}

func newLogger(id string,show bool) *Logger {
	l := new(Logger)
	l.sessionID = id
	l.show=show;
	return l
}

func (logger *Logger) Print(message interface{}) {
	if logger.show==true{
	log.Printf("%s   %s", logger.sessionID, message)
	}
}

func (logger *Logger) Printf(format string, v ...interface{}) {
	if logger.show==true {
		logger.Print(fmt.Sprintf(format, v...))
	}
}

func (logger *Logger) PrintCommand(command string, params string) {
	if logger.show==true {
		if command == "PASS" {
			log.Printf("%s > PASS ****", logger.sessionID)
		} else {
			log.Printf("%s > %s %s", logger.sessionID, command, params)
		}
	}
}

func (logger *Logger) PrintResponse(code int, message string) {
	if logger.show==true {
		log.Printf("%s < %d %s", logger.sessionID, code, message)
	}
}
