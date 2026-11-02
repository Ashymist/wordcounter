package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	res, err := count(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func count(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	wordcount := 0

	for scanner.Scan() {
		wordcount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordcount, nil
}
