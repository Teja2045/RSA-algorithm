package main

import (
	"fmt"
	src "rsa/src"
)

func testing_main() {
	fmt.Println(src.Encrypt("hi"))

	fmt.Println(string(src.Decrypt([]int{27800, 194474})))
}
