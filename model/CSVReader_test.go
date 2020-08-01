package model

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"
)

func TestCSVReader_removeHead(t *testing.T) {
	type args struct {
		n    int
		file *csv.Reader
	}
	f, err := os.Open("removeHeadTestData.csv")
	if err != nil {
		t.Errorf("removeHeadTestData.csv not found")
	}
	defer f.Close()
	r := csv.NewReader(f)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "removeHeadTestData.csv",
			args: args{
				n:    2,
				file: r,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSVReader{}
			if err := c.removeHead(tt.args.n, tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("removeHead() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCSVReader_parseCSV(t *testing.T) {
	type args struct {
		csv *csv.Reader
	}

	f, err := os.Open("removeHeadTestData.csv")
	if err != nil {
		t.Errorf("removeHeadTestData.csv not found")
	}
	defer f.Close()
	r := csv.NewReader(f)

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "ReadCSVTEST",
			args: args{
				csv: r,
			},
			want: [][]string{
				{"removeHead"},
				{"removeHead2"},
				{"Body", "Hoge"},
				{"Body2", "Fuga"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSVReader{}
			if got := c.parseCSV(tt.args.csv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
