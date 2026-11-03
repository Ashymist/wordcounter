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
	flag.Parse()

	res, err := count(os.Stdin, *lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func count(r io.Reader, lines bool) (int, error) {
	scanner := bufio.NewScanner(r)

	if !lines {
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
