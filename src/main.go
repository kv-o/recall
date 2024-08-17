package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

var (
	bflag bool
	lflag bool
	oflag bool
	sflag bool
	wflag string
)

func init() {
	flag.BoolVar(&bflag, "b", false, "query with back of the flashcard")
	flag.BoolVar(&lflag, "l", false, "use specified file as a Leitner learning box")
	flag.BoolVar(&oflag, "o", false, "overwrite specified file with wrongly answered terms")
	flag.BoolVar(&sflag, "s", false, "shuffle the flashcard deck")
	flag.StringVar(&wflag, "w", "", "wrongly answered terms output file")
}

func shuffle[T any](slice []T) {
	for i := range slice {
		randnum, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			fmt.Fprintln(os.Stderr, "memorize: cannot read from rand source")
			os.Exit(1)
		}
		j := int(randnum.Int64())
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func prompt(scanner *bufio.Scanner, term string) (wrong bool) {
	i, j := 0, 1
	if bflag {
		i, j = 1, 0
	}
	var side3 string
	pair := strings.Split(term, "\t")
	if len(pair) < 2 {
		fmt.Fprintf(os.Stderr, "memorize: malformed flashcard: %s\n", term)
		os.Exit(1)
	} else if len(pair) > 2 {
		side3 = pair[2]
	}
	fmt.Printf("%s: ", pair[i])
	scanner.Scan()
	msg := "Correct!"
	guess := scanner.Text()
	if guess != pair[j] {
		msg = pair[j]
		wrong = true
	}
	if side3 != "" {
		fmt.Printf("%s [%s]\n", msg, side3)
	} else {
		fmt.Println(msg)
	}
	return
}

func leitner() {
	fmt.Fprintln(os.Stderr, "memorize: leitner system is unimplemented")
	os.Exit(1)
}

func plain() {
	var wfile *os.File
	file, err := os.OpenFile(flag.Arg(0), os.O_RDWR, 0640)
	if err != nil {
		fmt.Fprintf(os.Stderr, "memorize: cannot open %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	defer file.Close()
	var terms []string
	lines := bufio.NewScanner(file)
	scanner := bufio.NewScanner(os.Stdin)
	if wflag != "" {
		wfile, err = os.OpenFile(wflag, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0640)
		if err != nil {
			fmt.Fprintf(os.Stderr, "memorize: cannot open %s: %v\n", wflag, err)
			os.Exit(1)
		}
		defer wfile.Close()
	}
	for lines.Scan() {
		if sflag || oflag {
			terms = append(terms, lines.Text())
		} else {
			wrong := prompt(scanner, lines.Text())
			if wrong && wflag != "" {
				fmt.Fprintln(wfile, lines.Text())
			}
		}
	}
	if sflag {
		shuffle(terms)
	}
	if oflag {
		wfile = file
		err = file.Truncate(0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "memorize: error truncating file")
			os.Exit(1)
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "memorize: error seeking start of file")
			os.Exit(1)
		}
	}
	if sflag || oflag {
		for _, term := range terms {
			wrong := prompt(scanner, term)
			if wrong && (wflag != "" || oflag) {
				fmt.Fprintln(wfile, term)
			}
		}
	}
	if err := lines.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "memorize: error reading %s: %v\n", flag.Arg(1), err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	lcond := lflag && (oflag || sflag || wflag != "")
	wcond := oflag && wflag != ""
	if len(flag.Args()) != 1 || lcond || wcond {
		fmt.Fprintln(os.Stderr, "memorize: invalid invocation")
		os.Exit(1)
	}
	if lflag {
		leitner()
	} else {
		plain()
	}
}
