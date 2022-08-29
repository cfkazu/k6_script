package main

import (
	"fmt"
	"os/exec"
	"unsafe"
)

func main() {
	//k6_only()
	k6_simple("test.k6.io")
}
func k6_only() {
	out, err := exec.Command("k6", "run", "k6_test.js").Output()
	print_bytes(out)
	fmt.Println(err)
}
func k6_simple(url string) {
	out, err := exec.Command("k6", "run", "-e", "MY_HOSTNAME="+url, "k6_env.js").Output()
	print_bytes(out)
	fmt.Println(err)
}
func print_bytes(b []byte) {
	fmt.Println(*(*string)(unsafe.Pointer(&b)))
}
