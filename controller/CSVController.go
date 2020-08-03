package controller

import (
	"kadai11/model"
	"kadai11/view"
	"log"
)

type CSVController struct {
	config    model.ImpConfig
	reader    model.ImpReader
	formatter Formatter
	writer    view.ImpWriter
}

func NewCSVController(config model.ImpConfig, reader model.CSVReader, formatter Formatter, writer view.ImpWriter) *CSVController {
	return &CSVController{
		config:    config,
		reader:    reader,
		formatter: formatter,
		writer:    writer,
	}
}

func (c *CSVController) Run() {
	rootDirs, inDirs := c.config.GetDirsNames()
	filenames := c.config.GetFilesName()

	c.writer.MakeDir("MyData")
	for i, dir := range rootDirs {
		c.writer.MakeDir("MyData/" + dir)
		for j := 0; j < len(inDirs[i]); j++ {
			c.writer.MakeDir("MyData/" + dir + "/" + inDirs[i][j])
			for _, filename := range filenames {
				log.Println("GetData:" + "RxData/" + dir + "/" + inDirs[i][j] + "/" + filename)
				h, _ := c.reader.GetData("RxData/" + dir + "/" + inDirs[i][j] + "/" + filename)
				fomattedData := c.formatter.FormatData(h)
				log.Println("WriteData:" + "MyData/" + dir + "/" + inDirs[i][j] + "/" + filename)
				c.writer.OverWriteFile("MyData/"+dir+"/"+inDirs[i][j]+"/"+filename, fomattedData)
			}
		}
	}
}
