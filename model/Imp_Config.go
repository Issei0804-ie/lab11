package model

type ImpConfig interface {
	GetDirsNames() ([]string, [][]string)
	GetFilesName() []string
}
