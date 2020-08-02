package model

type ImpReader interface {
	GetData(string) ([][]string, error)
}
