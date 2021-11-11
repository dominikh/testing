package pkg

import "fmt"

func Fn() {
	var s string
	//lint:ignore SA1006 this is fine
	fmt.Printf(s)
}
