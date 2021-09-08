package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	s := Add(2, 1)
	if s != 3 {
		log.Fatalf("expect 3,but got %d", s)
	}
	s = Add(-1, 1)
	if s != 0 {
		log.Fatalf("expect 0,but got %d", s)
	}
}
