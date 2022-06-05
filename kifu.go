package kifu

import (
	"fmt"
	"log"
)

const (
	ERROR = "ERRR"
	FATAL = "FATL"
	INFO  = "INFO"
	WARN  = "WARN"
)

func Info(identifiers []string, content interface{}) {
	logText := constructLogText(INFO, identifiers, content)
	log.Println(logText)
}

func Warn(identifiers []string, content interface{}) {
	logText := constructLogText(INFO, identifiers, content)
	log.Println(logText)
}

func Error(identifiers []string, content interface{}) {
	logText := constructLogText(ERROR, identifiers, content)
	log.Println(logText)
}

func Fatal(identifiers []string, content interface{}) {
	logText := constructLogText(FATAL, identifiers, content)
	log.Fatalln(logText)
}

func constructLogText(logType string, identifiers []string, content interface{}) string {
	logText := logType

	for _, identifier := range identifiers {
		logText += fmt.Sprintf(" [%v]", identifier)
	}

	if content != nil {
		logText += fmt.Sprintf(" %v", content)
	}

	return logText
}
