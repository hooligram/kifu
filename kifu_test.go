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
	// Output: ╮(︶︿︶)╭ · Netflix
}

func ExampleError() {
	kifu.Error("%v", "Facebook")
	// Output: (〃＞＿＜;〃) · Facebook
}

func ExampleFatal() {
	defer func() {
		recover()
	}()
	kifu.Fatal("%v", "Amazon")
	// Output: (ｏ・_・)ノ”(ノ_<、) · Amazon
}
