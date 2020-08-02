package main

import (
	"kadai11/controller"
	"kadai11/model"
	"kadai11/view"
)

func main() {
	c := model.NewConfig()
	r := model.NewCSVReader()
	f := controller.NewFormatter()
	w :=view.NewCSVWriter()
	cont := controller.NewCSVController(c, *r, *f, w)
	cont.Run()
}
