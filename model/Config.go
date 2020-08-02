package model

import (
	"fmt"
	"strconv"
)

type Config struct {
	Year     int
	Month    []int
	Filename []string
}

func NewConfig() *Config {
	return &Config{
		Year:     2009,
		Month:    []int{6, 10, 11, 12},
		Filename: []string{"192.168.100.11_csv.log", "192.168.100.9_csv.log"},
	}
}

func (c *Config) EndOfTheMonthIs30(month int) bool {
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return true
	} else {
		return false
	}
}

func (c *Config) GetDirsNames() ([]string, [][]string) {
	var dirNames [][]string
	yearMonth := c.GetYearMonth()
	for j, month := range c.Month {
		counter := 31
		if c.EndOfTheMonthIs30(month) {
			counter = 30
		}
		var dirName []string
		for i := 1; i < counter+1; i++ {
			dirName = append(dirName, yearMonth[j]+fmt.Sprintf("%02d", i))
		}
		dirNames = append(dirNames, dirName)
	}
	return yearMonth, dirNames
}

func (c *Config) GetFilesName() []string {
	return c.Filename
}

func (c *Config) GetYearMonth() []string {
	var YearMonth []string

	for _, month := range c.Month {
		// 0 padding
		YearMonth = append(YearMonth, strconv.Itoa(c.Year)+fmt.Sprintf("%02d", month))
	}
	return YearMonth
}
