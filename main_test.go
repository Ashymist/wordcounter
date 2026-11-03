package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	expectedCount := 4

	res, err := count(b, false)

	if res != expectedCount || err != nil {
		t.Errorf("Expected %v,nil; got \"%v\", \"%v\"\n", expectedCount, res, err)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1\n\n\nword2\nword4\n")

	expectedLines := 5

	res, err := count(b, true)

	if res != expectedLines || err != nil {
		t.Errorf("Expected %v,nil; got \"%v\", \"%v\"\n", expectedLines, res, err)
	}
}
