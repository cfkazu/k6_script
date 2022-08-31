package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	tags := ""
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), ",")
		urls += str[0] + ";"
		tags += str[1] + ";"
	}
	fmt.Println(urls)
	fmt.Println(tags)
	out, err := exec.Command("k6", "run", "-e", "HOSTNAMES="+urls, "-e", "TAGS="+tags, "k6_env_multiple.js", "-o", "json=output.json").Output()

	print_bytes(out)
	fmt.Println(err)
}
func print_bytes(b []byte) {
	fmt.Println(*(*string)(unsafe.Pointer(&b)))
}
