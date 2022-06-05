package kifu

import (
	"fmt"
	"os"
	"strings"
)

func Info(format string, a ...any) {
	fmt.Printf("٩(｡•́‿•̀｡)۶ · %v\n", fmt.Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	fmt.Printf("(；⌣̀_⌣́) · %v\n", fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	fmt.Printf("(╬ Ò﹏Ó) · %v\n", fmt.Sprintf(format, a...))
}

func Fatal(format string, a ...any) {
	fmt.Printf("凸(￣ヘ￣) · %v\n", fmt.Sprintf(format, a...))
	if strings.HasSuffix(os.Args[0], "test") {
		return
	}
	os.Exit(1)
}
