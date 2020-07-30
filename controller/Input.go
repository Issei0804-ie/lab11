package controller

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

type Input struct {
	data [][]string
	conf config
}

func NewInput() *Input {
	return &Input{
		conf: newConfig(),
	}
}

func (in *Input) ReadCSV(filepath string) [][]string {
	f := in.readFile(filepath)
	r := csv.NewReader(f)

	if err := in.removeHead(2, r); err != nil {
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

	return nil
}

//読み込んだファイルの先頭2行を破棄します．
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
	DirConf  dirConf
	FileConf fileConf
}

func newConfig() config {
	return config{
		DirConf:  dirConf{},
		FileConf: fileConf{},
	}
}

type dirConf struct {
	Year  int
	Month []int
}

func newDirConf() dirConf {
	return dirConf{
		Year:  2009,
		Month: []int{6, 10, 11, 12},
	}
}

func (dc *dirConf) EndOfTheMonthIs30(month int) bool {
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return true
	} else {
		return false
	}
}

func newFileConf() fileConf {
	return fileConf{
		filename: []string{"192.168.100.11_csv.log", "192.168.100.9_csv.log"},
	}
}

type fileConf struct {
	filename []string
}
