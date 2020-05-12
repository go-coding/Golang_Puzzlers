package main

import "log"

type Params struct {
	a, b, c int
}

func doIt(p Params) int {
	return p.a + p.b + p.c
}

func main() {
	// you can call it without specifying all parameters
	res0 := doIt(Params{a: 1, c: 9})

	log.Printf("res0 => %v ", res0)

	res1 := doIt(Params{a: 1, c: 9, b: 10})

	log.Printf("res1 => %v ", res1)

	res2 := doIt(Params{a: 1, c: 9, b: 1})

	log.Printf("res1 => %v ", res2)

}
