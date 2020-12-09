package main

import "fmt"

type DataFrame struct {
	Data []Series
}

type Series interface {
	Len() int
	Type() string
}
