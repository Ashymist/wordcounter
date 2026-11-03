package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	expectedCount := 4

	res, err := count(b, false, false)

	if res != expectedCount || err != nil {
		t.Errorf("Expected %v,nil; got \"%v\", \"%v\"\n", expectedCount, res, err)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1\n\n\nword2\nword4\n")

	expectedLines := 5

	res, err := count(b, true, false)

	if res != expectedLines || err != nil {
		t.Errorf("Expected %v,nil; got \"%v\", \"%v\"\n", expectedLines, res, err)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("wordword")

	expectedBytes := 8

	res, err := count(b, false, true)

	if res != expectedBytes || err != nil {
		t.Errorf("Expected %v,nil; got \"%v\", \"%v\"\n", expectedBytes, res, err)
	}
}

func TestMultipleFlagsError(t *testing.T) {
	b := bytes.NewBufferString("wordword")

	res, err := count(b, true, true)

	if err == nil && !errors.Is(err, MultipleFlagError) {
		t.Errorf("Expected 0, %v; got \"%v\", \"%v\"\n", MultipleFlagError, res, err)
	}
}
