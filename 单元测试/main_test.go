package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	if sum == 5 {
		t.Log("result is ok")
	} else {
		t.Fatal("result is wrong")
	}
}

func TestFib(t *testing.T) {
	result := Fib(10)
	t.Log("result is :", result)
}
