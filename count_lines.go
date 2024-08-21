package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func print_line_count_from_stdin() {

	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		fmt.Println("scanned", line)
		counts[line]++
	}

	for line, count := range counts {
		fmt.Printf("%s\t:%d\n", line, count)
		if count > 1 {
			fmt.Printf("%s\t:%d\n", line, count)
		}
	}

}

func print_line_count_and_occuring_files() {
	lineCount := make(map[string]int)
	lineFile := make(map[string]string)
	files := os.Args[1:]
	for _, fName := range files {
		f, _ := os.Open(fName)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			lineCount[line]++
			if !strings.Contains(lineFile[line], fName) {
				lineFile[line] += fName
			}
		}
	}

	for line, n := range lineCount {
		fmt.Printf("%s\t%d\t%s", line, n, lineFile[line])
	}
}
