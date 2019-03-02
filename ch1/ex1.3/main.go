// Echo prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	measureForLoop()
	measureRange()
	measureJoin()
}

func measureJoin() {
	start := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Printf("strings.Join: %d Nanoseconds\n\n", time.Since(start).Nanoseconds())
}

func measureForLoop() {

	start := time.Now()
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Printf("For Loop: %d Nanoseconds\n\n", time.Since(start).Nanoseconds())
}

func measureRange() {
	start := time.Now()
	var s, sep string

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Printf("Range: %d Nanoseconds\n\n", time.Since(start).Nanoseconds())
}
