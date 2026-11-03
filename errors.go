package main

import "errors"

var MultipleFlagError = errors.New("Only one counting flag can be selected")
