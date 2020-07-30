package view

type Output struct {
	ImpFormatter
}

type ImpFormatter interface {
	GetData() [][]string
}

func (out *Output) MakeDir() {

}
func (out *Output) OverWriteToFile() {

}
