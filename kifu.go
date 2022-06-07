package kifu

import (
	"fmt"
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
	msg := fmt.Sprintf("(ｏ・_・)ノ”(ノ_<、) · %v\n", fmt.Sprintf(format, a...))
	fmt.Print(msg)
	panic(msg)
}
