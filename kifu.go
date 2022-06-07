package kifu

import (
	"fmt"
	"os"
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
	os.Exit(1)
}
