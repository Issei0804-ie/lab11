package view

type ImpWriter interface {
	OverWriteFile(filepath string, data [][]string)
	MakeDir(dirpath string)
}
