package model

type Formatter struct {
	ImpInput
}

type ImpInput interface {
	GetData() [][]string
}
