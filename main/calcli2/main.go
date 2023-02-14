package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"calcy"
)

func main() {
	var op string
	flag.StringVar(&op, "op", "+", "Pick one: + - * / ?")
	flag.Parse()

	handler := NewCLIHandler(calculators[op])

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

type CLIHandler struct {
	calc calcy.Calculator
}

func NewCLIHandler(calculator calcy.Calculator) *CLIHandler {
	return &CLIHandler{calc: calculator}
}

func (this *CLIHandler) Handle(args []string) error {
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	b, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	fmt.Println(this.calc.Calculate(a, b))
	return nil
}
