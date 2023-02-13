package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"calcy"
)

func main() {
	http.Handle("/add", NewHandler(calcy.Addition{}))
	http.Handle("/sub", NewHandler(calcy.Subtraction{}))
	http.Handle("/mul", NewHandler(calcy.Multiplication{}))
	http.Handle("/div", NewHandler(calcy.Division{}))
	http.Handle("/bog", NewHandler(calcy.Bogus{}))

	log.Println("Listening on localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

type Handler struct {
	calculator calcy.Calculator
}

func NewHandler(calculator calcy.Calculator) http.Handler {
	return &Handler{calculator: calculator}
}

func (this *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	rawA := query.Get("a")
	a, err := strconv.Atoi(rawA)
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = fmt.Fprintf(writer, "invalid 'a' parameter: [%s]", rawA)
		return
	}

	rawB := query.Get("b")
	b, err := strconv.Atoi(rawB)
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = fmt.Fprintf(writer, "invalid 'b' parameter: [%s]", rawB)
		return
	}

	writer.WriteHeader(200)
	_, err = fmt.Fprintln(writer, this.calculator.Calculate(a, b))
	if err != nil {
		log.Println("Failed to write response:", err)
	}
}
