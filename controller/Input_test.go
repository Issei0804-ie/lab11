package controller

import (
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
