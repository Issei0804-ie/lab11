package controller

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Input struct {
	data [][]string
	conf config
}

func NewInput() *Input {
	in := &Input{
		conf: newConfig(),
	}
	in.Run()
	return in
}

func (in *Input) Run() {
	//yearMonth := in.GetYearMonth()

}

func (in *Input) GetDirNames() [][]string {
	var dirNames [][]string
	yearMonth := in.GetYearMonth()
	for j, month := range in.conf.DirConf.Month {
		counter := 31
		if in.conf.DirConf.EndOfTheMonthIs30(month) {
			counter = 30
		}
		var dirName []string
		for i := 1; i < counter+1; i++ {
			dirName = append(dirName, yearMonth[j]+fmt.Sprintf("%02d", i))
		}
		dirNames = append(dirNames, dirName)
	}
	return dirNames
}

func (in *Input) GetYearMonth() []string {
	var YearMonth []string

	for _, month := range in.conf.DirConf.Month {
		// 0 padding
		YearMonth = append(YearMonth, strconv.Itoa(in.conf.DirConf.Year)+fmt.Sprintf("%02d", month))
	}
	return YearMonth
}

//指定されたcsvファイルを読み取ります．
func (in *Input) ReadCSV(filepath string) [][]string {
	f := in.readFile(filepath)
	r := csv.NewReader(f)

	row := 2
	if err := in.removeHead(row, r); err != nil {
		fmt.Println(err.Error())
	}
	return in.parseCSV(r)
}

//指定されたファイルパスのファイルを読み込みます
func (in *Input) readFile(filepath string) *os.File {
	//指定したファイルが存在しなければ
	_, err := os.Stat(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

func (in *Input) parseCSV(r *csv.Reader) [][]string {
	var records [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		records = append(records, record)
	}
	return records
}

//読み込んだファイルの先頭n行を破棄します．
func (in *Input) removeHead(n int, r *csv.Reader) error {
	for i := 0; i < 2; i++ {
		_, err := r.Read()
		if err == io.EOF {
			return errors.New("file exists but not wrote")
		}
	}
	return nil
}

func (in *Input) GetData() [][]string {
	return nil
}

type config struct {
	DirConf  DirConf
	FileConf FileConf
}

func newConfig() config {
	return config{
		DirConf:  newDirConf(),
		FileConf: newFileConf(),
	}
}

type DirConf struct {
	Year  int
	Month []int
}

func newDirConf() DirConf {
	return DirConf{
		Year:  2009,
		Month: []int{6, 10, 11, 12},
	}
}

func (dc *DirConf) EndOfTheMonthIs30(month int) bool {
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return true
	} else {
		return false
	}
}

func newFileConf() FileConf {
	return FileConf{
		Filename: []string{"192.168.100.11_csv.log", "192.168.100.9_csv.log"},
	}
}

type FileConf struct {
	Filename []string
}
