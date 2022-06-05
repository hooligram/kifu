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

func Info(format string, a ...any) {
	log.Printf("%v %v\n", INFO, fmt.Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	log.Printf("%v %v\n", WARN, fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	log.Printf("%v %v\n", ERROR, fmt.Sprintf(format, a...))
}

func Fatal(format string, a ...any) {
	log.Fatalf("%v %v\n", FATAL, fmt.Sprintf(format, a...))
}
