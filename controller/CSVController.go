package controller

import (
	"kadai11/model"
	"kadai11/view"
)

type CSVController struct {
	config    model.ImpConfig
	reader    model.ImpReader
	formatter Formatter
	writer view.ImpWriter
}

func NewCSVController(config model.ImpConfig, reader model.CSVReader, formatter Formatter, writer view.ImpWriter) *CSVController {
	return &CSVController{
		config:    config,
		reader:    reader,
		formatter: formatter,
		writer: writer,
	}
}

func (c *CSVController) Run() {
	a, b :=c.config.GetDirsNames()
	filenames := c.config.GetFilesName()

	h, _ :=c.reader.GetData("RxData/"+a[0] + "/" +b[0][0] + "/" + filenames[0])
	fomattedData :=c.formatter.FormatData(h)

	c.writer.MakeDir("MyData/" + a[0])
	c.writer.MakeDir("MyData/" + a[0] + "/" +b[0][0])
	c.writer.OverWriteFile("MyData/"+a[0] + "/" +b[0][0] + "/" + filenames[0], fomattedData)
}
