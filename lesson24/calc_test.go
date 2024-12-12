package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	if res := Add(10, 20); res != 30 {
		log.Fatalf("Add(10, 20) fail. Expected %d, got %d\n", 30, res)
	}

	if res := Add(30, 30); res != 60 {
		log.Fatalf("Add(30, 30) fail. Expected %d, got %d\n", 60, res)
	}
}

func TestSub(t *testing.T) {

}

func TestMult(t *testing.T) {

}
