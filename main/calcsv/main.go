package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"calcy"
)

func main() {
	reader := csv.NewReader(os.Stdin)
	writer := csv.NewWriter(os.Stdout)

	handler := NewCSVHandler(reader, writer)

	err := handler.Handle()
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

type CSVHandler struct {
	reader *csv.Reader
	writer *csv.Writer
}

func NewCSVHandler(reader *csv.Reader, writer *csv.Writer) *CSVHandler {
	return &CSVHandler{reader: reader, writer: writer}
}

func (this *CSVHandler) Handle() error {
	defer this.writer.Flush()
	for {
		record, err := this.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		a, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}

		b, err := strconv.Atoi(record[2])
		if err != nil {
			return err
		}

		calculator, ok := calculators[record[1]]
		if !ok {
			return fmt.Errorf("unsupported operation: %s", record[1])
		}

		result := calculator.Calculate(a, b)

		err = this.writer.Write(append(record, fmt.Sprint(result)))
		if err != nil {
			return err
		}
	}
	return nil
}
