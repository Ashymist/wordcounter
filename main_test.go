package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	expectedCount := 4

	res, err := count(b)

	if res != expectedCount || err != nil {
		t.Errorf("Expected %v,nil; got \"%v\", \"%v\"\n", expectedCount, res, err)
	}
}
