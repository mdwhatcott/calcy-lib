package main

import (
	"flag"
	"log"
	"os"

	"calcy"
	"calcy/handlers"
)

func main() {
	var op string
	flag.StringVar(&op, "op", "+", "Pick one: + - * / ?")
	flag.Parse()

	handler := handlers.NewCLIHandler(calculators[op], os.Stdout)

	err := handler.Handle(flag.Args())
	if err != nil {
		log.Fatalln(err)
	}
}

var calculators = map[string]calcy.Calculator{
	"+": calcy.Addition{},
	"-": calcy.Subtraction{},
	"*": calcy.Multiplication{},
	"/": calcy.Division{},
	"?": calcy.Bogus{Offset: 42},
}
