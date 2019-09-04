package main

import (
	"fmt"
	"github.com/wizcas/arse/parser"
	"log"
)

func main() {
	println("Hello my arse!")
	data, err := parser.Load("./tests/test_config.yml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Loaded ArseFile:\n%v\n", data)
}
