package kifu

import (
	"fmt"
	"log"
)

// Tags to more easily identify log types.
const (
	BODY  = "BODY"
	CLOSE = "CLSE"
	INFO  = "INFO"
	OPEN  = "OPEN"
)

// Body .
func Body(identifiers []string, content interface{}) {
	constructLogText(BODY, identifiers, content)
}

// Close .
func Close(identifiers []string, content interface{}) {
	constructLogText(CLOSE, identifiers, content)
}

// Info .
func Info(identifiers []string, content interface{}) {
	constructLogText(INFO, identifiers, content)
}

// Open .
func Open(identifiers []string, content interface{}) {
	constructLogText(OPEN, identifiers, content)
}

func constructLogText(logType string, identifiers []string, content interface{}) {
	logText := logType

	for _, identifier := range identifiers {
		logText += fmt.Sprintf(" [%v]", identifier)
	}

	if content != nil {
		logText += fmt.Sprintf(" %v", content)
	}

	log.Println(logText)
}
