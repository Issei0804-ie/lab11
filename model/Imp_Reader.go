package model

type ImpReader interface {
	GetData(filepath string) [][]string
}
