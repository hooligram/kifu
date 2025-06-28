package kifu

import (
	"fmt"
	"log"
)

func Info(format string, a ...any) {
	log.Printf("INFO · %v", fmt.Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	log.Printf("WARN · %v", fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	log.Printf("ERRO · %v", fmt.Sprintf(format, a...))
}

func Fatal(format string, a ...any) {
	log.Fatalf("FATL · %v\n", fmt.Sprintf(format, a...))
}
