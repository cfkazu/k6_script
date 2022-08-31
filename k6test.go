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
	k6_filename(os.Args[1])
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
	js_filename := get_jsfile_from_options(os.Args[2])
	fmt.Println(urls)
	fmt.Println(tags)
	out, err := exec.Command("k6", "run", "-e", "HOSTNAMES="+urls, "-e", "TAGS="+tags, js_filename, "-o", "json=output.json").Output()

	print_bytes(out)
	fmt.Println(err)
}
func get_jsfile_from_options(options string) string {
	switch options {
	case "simple":
		return "k6_test.js"
	case "stress":
		return "k6_stress.js"
	case "spike":
		return "k6_spike.js"
	}

	return "k6_" + options + ".js"
}
func print_bytes(b []byte) {
	fmt.Println(*(*string)(unsafe.Pointer(&b)))
}
