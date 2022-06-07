package kifu_test

import (
	"github.com/hooligram/kifu"
)

func ExampleInfo() {
	kifu.Info("%v %v %v", "Google", "Microsoft", "Apple")
	// Output: ٩(｡•́‿•̀｡)۶ · Google Microsoft Apple
}

func ExampleWarn() {
	kifu.Warn("%v", "Netflix")
	// Output: (；⌣̀_⌣́) · Netflix
}

func ExampleError() {
	kifu.Error("%v", "Facebook")
	// Output: (╬ Ò﹏Ó) · Facebook
}
