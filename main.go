package main

import (
	"io/ioutil"
	"kadai11/controller"
	"kadai11/model"
	"kadai11/view"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// logの非表示
	log.SetOutput(ioutil.Discard)
	c := model.NewConfig()
	r := model.NewCSVReader()
	f := controller.NewFormatter()
	w := view.NewCSVWriter()
	cont := controller.NewCSVController(c, *r, *f, w)
	cont.Run()

	// logの表示
	log.SetOutput(os.Stderr)
	end := time.Now()

	log.Printf("%v秒\n", (end.Sub(start)).Seconds())
}
