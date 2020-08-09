package model

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path/filepath"
)

type GraphConfig struct {
	FilepathRainData []string
}

func NewGraphConfig() GraphConfig {
	g := GraphConfig{}
	g.FilepathRainData = g.getGraphConfig()
	return g
}

func (g *GraphConfig) getGraphConfig() []string {
	pattern := ".RainData/*/*.csv"
	files, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}
	return files
}

func (g *GraphConfig) GetRainData(filepath string)[][]string{
	f, _ := os.Open(filepath)
	r := csv.NewReader(f)

	g.removeHead(2,r)

	data := g.parseCSV(r)

	return data
}

func (c *GraphConfig) parseCSV(csv *csv.Reader) [][]string {
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

func (c *GraphConfig) removeHead(n int, csv *csv.Reader) error {
	for i := 0; i < 2; i++ {
		_, err := csv.Read()
		if err == io.EOF {
			return errors.New("file exists but not wrote")
		}
	}
	return nil
}
