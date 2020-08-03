package view

import (
	"encoding/csv"
	"os"
)

type CSVWriter struct {
}

func NewCSVWriter() *CSVWriter {
	return &CSVWriter{}
}

func (c CSVWriter) OverWriteFile(filepath string, data [][]string) {
	_, err := os.Stat(filepath)
	if err != nil {
		f, _ := os.Create(filepath)
		w := csv.NewWriter(f)
		w.WriteAll(data)
	}
}

func (c CSVWriter) MakeDir(dirpath string) {
	if !c.CheckExistsDir(dirpath) {
		err := os.Mkdir(dirpath, 0777)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (c CSVWriter) CheckExistsDir(dirpath string) bool {
	_, err := os.Stat(dirpath)
	if err != nil {
		return false
	}
	return true
}
