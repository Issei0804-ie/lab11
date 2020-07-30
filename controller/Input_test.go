package controller

import (
	"encoding/csv"
	"os"
	"testing"
)

func Test_dirConf_EndOfTheMonthIs30(t *testing.T) {
	type fields struct {
		Year  int
		Month []int
	}
	type args struct {
		month int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "Jone",
			fields: fields{
				Year:  2009,
				Month: []int{6, 10, 11, 12},
			},
			args: args{
				month: 6,
			},
			want: true,
		},
		{
			name: "October",
			fields: fields{
				Year:  2009,
				Month: []int{6, 10, 11, 12},
			},
			args: args{
				month: 10,
			},
			want: false,
		},
		{
			name: "November",
			fields: fields{
				Year:  2009,
				Month: []int{6, 10, 11, 12},
			},
			args: args{
				month: 11,
			},
			want: true,
		},
		{
			name: "December",
			fields: fields{
				Year:  2009,
				Month: []int{6, 10, 11, 12},
			},
			args: args{
				month: 12,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &dirConf{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
			}
			if got := dc.EndOfTheMonthIs30(tt.args.month); got != tt.want {
				t.Errorf("EndOfTheMonthIs30() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInput_removeHead(t *testing.T) {
	type fields struct {
		data [][]string
		conf config
	}
	type args struct {
		n int
		r *csv.Reader
	}

	f, err := os.Open("removeHeadTestData.csv")
	if err != nil {
		t.Errorf("removeHeadTestData.csv not found")
	}
	defer f.Close()

	r := csv.NewReader(f)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "removeHeadTestData.csv",
			fields: fields{
				data: [][]string{},
				conf: config{},
			},
			args: args{
				n: 2,
				r: r,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := &Input{
				data: tt.fields.data,
				conf: tt.fields.conf,
			}
			if err := in.removeHead(tt.args.n, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("removeHead() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
