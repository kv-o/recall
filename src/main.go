package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// add option to output incorrect terms to a certain file

var (
	rflag bool
)

func init() {
	flag.BoolVar(&rflag, "r", false, "query with reverse of the flashcard")
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "memorize: invalid invocation")
		os.Exit(1)
	}
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "memorize: cannot open %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	i := 0
	j := 1
	var guess string
	terms := bufio.NewScanner(file)
	scanner := bufio.NewScanner(os.Stdin)
	if rflag {
		i = 1
		j = 0
	}
	for terms.Scan() {
		pair := strings.Split(terms.Text(), "\t")
		fmt.Printf("%s: ", pair[i])
		scanner.Scan()
		guess = scanner.Text()
		if guess == pair[j] {
			fmt.Println("Correct!")
		} else {
			fmt.Println(pair[j])
		}
	}
	if err := terms.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "memorize: error reading %s: %v", flag.Arg(1), err)
		os.Exit(1)
	}
}
