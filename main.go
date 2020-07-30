package main

import "kadai11/parser"

func main() {
	p, err := parser.NewParser("200906/20090601/192.168.100.11_csv.log")
	if err != nil {
		print(err.Error())
	}
	p.MergeRecords()
}
