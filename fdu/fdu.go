package fdu

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 reads input from standard input and prints the count and text of lines that appear more than once.
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {

			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// I'm trying to make Dup2 function more expressive with functional approach
func Dup2V2(files []string) {
	if len(files) == 0 {
		fmt.Println("Belum ada file yang diinput")
	}

	result := make(map[string]int)

	for _, file := range files {
		countLinesV2(file, result)
	}

	for line, n := range result {
		if n > 1 {
			fmt.Printf("Kata \"%s\" terulang sebanyak %d kali\n", line, n)
		}
	}

}

func countLinesV2(file string, counts map[string]int) {
	content, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	}
	input := bufio.NewScanner(content)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// countLines(os.Stdin, counts)
		fmt.Println("Please provide file names")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Kata \"%s\" terulang sebanyak %d kali\n", line, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
