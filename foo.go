package pkg

import "regexp"

func fn() {
	regexp.MustCompile("+")
	regexp.MustCompile("+")
}
