package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	countBytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	res, err := count(os.Stdin, *lines, *countBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func count(r io.Reader, lines bool, countBytes bool) (int, error) {
	scanner := bufio.NewScanner(r)

	if lines && countBytes {
		return 0, MultipleFlagError
	}

	switch {
	case lines:
		scanner.Split(bufio.ScanLines)
	case countBytes:
		scanner.Split(bufio.ScanBytes)
	default:
		scanner.Split(bufio.ScanWords)
	}

	count := 0

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
