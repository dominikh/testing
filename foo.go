package pkg

import "fmt"
import . "io"

var _ = io.Reader

func Fn() {
	s := ""
	fmt.Printf(s)
}

func Fn2(x *int) {
	if x == nil {
		println()
	}
	_ = *x
}

func unused() {
}
