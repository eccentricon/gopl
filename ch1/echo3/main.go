package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(Echo(os.Args[1:]))
}
func Echo(args []string) string {
	return strings.Join(args, " ")
}
