package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// add option to output incorrect terms to a certain file

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "memorize: invalid invocation")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "memorize: cannot open %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	var guess string
	terms := bufio.NewScanner(file)
	scanner := bufio.NewScanner(os.Stdin)
	for terms.Scan() {
		pair := strings.Split(terms.Text(), "\t")
		fmt.Printf("%s: ", pair[0])
		scanner.Scan()
		guess = scanner.Text()
		if guess == pair[1] {
			fmt.Println("Correct!")
		} else {
			fmt.Println(pair[1])
		}
	}
	if err := terms.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "memorize: error reading %s: %v", os.Args[1], err)
		os.Exit(1)
	}
}
