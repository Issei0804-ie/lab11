package controller

import "os"

type Input struct {
	data [][]string
	conf config
}

func NewInput() *Input {
	return &Input{}
}

func (in *Input) ReadCSV() [][]string {
	return nil
}

func (in *Input) readFile(filepath string) *os.File {
	return nil
}

func (in *Input) parseCSV() [][]string {
	return nil
}

func (in *Input) removeHead(n int) {

}

func (in *Input) GetData() [][]string {
	return nil
}

type config struct {
	DirConf  dirConf
	FileConf fileConf
}

func newConfig() config{
	return config{
		DirConf:  dirConf{},
		FileConf: fileConf{},
	}
}

type dirConf struct {
	Year  int
	Month []int
}

func newDirConf() dirConf{
	return dirConf{
		Year:  2009,
		Month: []int{6,10,11,12},
	}
}

func (dc *dirConf) EndOfTheMonthIs30(month int) bool {
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return true
	} else {
		return false
	}
}

func newFileConf() fileConf{
	return fileConf{
		filename: []string{"192.168.100.11_csv.log","192.168.100.9_csv.log"},
	}
}

type fileConf struct {
	filename []string
}
