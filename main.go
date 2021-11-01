package main

import (
	"fmt"
	"log"

	bigint "github.com/abbos-ron2/go/bigint/bigint"
)

func main() {
	a, err :=
		bigint.NewInt("1")
	if err != nil {
		log.Println(err)
	}
	b, err := bigint.NewInt("100")
	if err != nil {
		log.Println(err)
	}

	c := bigint.Sub(a, b)
	fmt.Println(c)
}
