// Ex1.4 prints the count and text of lines that appear more than once
// in the input. In addition it prints the names of all files in which each
// duplicated lines occur. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	flines := make(map[string]map[string]bool)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, flines)
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts, flines)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			filenames := make([]string, len(flines[line]))
			i := 0
			for filename := range flines[line] {
				filenames[i] = filename
				i++
			}
			fmt.Printf("%s\t%d\t%s\n", line, n, strings.Join(filenames, " "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, flines map[string]map[string]bool) {

	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
		if nil == flines[input.Text()] {
			flines[input.Text()] = map[string]bool{}
		}
		flines[input.Text()][f.Name()] = true
	}
}
