package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"calcy"
)

func main() {
	var calculator calcy.Addition
	handler := NewCLIHandler(calculator)
	err := handler.Handle(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

type CLIHandler struct {
	calc calcy.Calculator
}

func NewCLIHandler(calculator calcy.Calculator) *CLIHandler {
	return &CLIHandler{calc: calculator}
}

func (this *CLIHandler) Handle(args []string) error {
	a, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	b, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}

	fmt.Println(this.calc.Calculate(a, b))
	return nil
}
