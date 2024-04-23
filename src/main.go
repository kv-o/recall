package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	bflag bool
	wflag string
)

// add -o flag to overwrite given file with wrongly
// answered terms, instead of outputting to wfile
func init() {
	flag.BoolVar(&bflag, "b", false, "query with reverse of the flashcard")
	flag.StringVar(&wflag, "w", "", "wrongly answered terms output file")
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
	if bflag {
		i = 1
		j = 0
	}
	wfile, err := os.OpenFile(wflag, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0640)
	if err != nil {
		fmt.Fprintf(os.Stderr, "memorize: cannot open %s: %v", wflag, err)
		os.Exit(1)
	}
	defer wfile.Close()
	for terms.Scan() {
		pair := strings.Split(terms.Text(), "\t")
		fmt.Printf("%s: ", pair[i])
		scanner.Scan()
		guess = scanner.Text()
		if guess == pair[j] {
			fmt.Println("Correct!")
		} else {
			fmt.Println(pair[j])
			fmt.Fprintln(wfile, terms.Text())
		}
	}
	if err := terms.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "memorize: error reading %s: %v", flag.Arg(1), err)
		os.Exit(1)
	}
	if wflag == "" {
		os.Exit(0)
	}
}
