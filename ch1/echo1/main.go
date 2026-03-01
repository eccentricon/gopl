// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Echo(os.Args[1:]))
}

func Echo(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}
