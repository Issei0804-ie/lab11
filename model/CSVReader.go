package model

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type CSVReader struct {
}

func NewCSVReader() *CSVReader {
	return &CSVReader{}
}

func (c CSVReader) GetData(filepath string) ([][]string, error) {
	log.Println("ReadFile....")
	f, err := c.ReadFile(filepath)
	log.Println("Finish")
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)

	row := 2
	if err := c.removeHead(row, r); err != nil {
		fmt.Println(err.Error())
	}

	log.Println("ParseToCSV....")
	defer log.Println("Finish")
	return c.parseCSV(r), nil
}

func (c *CSVReader) ReadFile(filepath string) (*os.File, error) {
	//指定したファイルが存在しなければ
	_, err := os.Stat(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	return f, nil
}

func (c *CSVReader) parseCSV(csv *csv.Reader) [][]string {
	var records [][]string
	for {
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		records = append(records, record)
	}
	return records
}

func (c *CSVReader) removeHead(n int, csv *csv.Reader) error {
	for i := 0; i < 2; i++ {
		_, err := csv.Read()
		if err == io.EOF {
			return errors.New("file exists but not wrote")
		}
	}
	return nil
}
