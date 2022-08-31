package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"unsafe"
)

func main() {
	//k6_only()
	//k6_simple("test.k6.io")
	k6_filename("hosts.txt")
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

func k6_filename(filename string) {
	data, err := os.Open(filename)
	fmt.Println(err)
	defer data.Close()
	scanner := bufio.NewScanner(data)
	urls := ""
	for scanner.Scan() {
		urls += scanner.Text() + ","
	}
	fmt.Println(urls)
	out, err := exec.Command("k6", "run", "-e", "HOSTNAMES="+urls, "k6_env_multiple.js", "-o", "json=output.json").Output()

	print_bytes(out)
	fmt.Println(err)
}
func print_bytes(b []byte) {
	fmt.Println(*(*string)(unsafe.Pointer(&b)))
}
