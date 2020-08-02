package controller

import (
	"reflect"
	"testing"
	"time"
)

func TestFormatter_GetAve(t *testing.T) {
	type fields struct {
		Average             int
		AveragePreviousFile int
	}
	type args struct {
		data     [][]string
		baseTime *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
		want2  int
	}{
		// TODO: Add test cases.
		{
			name: "not exception test",
			fields: fields{
				Average:             0,
				AveragePreviousFile: 0,
			},
			args: args{
				data: [][]string{
					{"00:00:02", "-5"},
					{"00:00:05", "-7"},
				},
				baseTime: testFuncDataTime("00:00:00"),
			},
			want:  []string{"00:00:00", "-6"},
			want2: 2,
		},
		{
			name:   "overflow data",
			fields: fields{},
			args: args{
				data: [][]string{
					{"00:00:02", "-5"},
					{"00:00:05", "-7"},
					{"00:00:07", "-8"},
					{"00:00:08", "-12"},
					{"00:00:12", "-7"},
				},
				baseTime: testFuncDataTime("00:00:00"),
			},
			want:  []string{"00:00:00", "-8"},
			want2: 4,
		},
		{
			name: "NoAveData",
			fields: fields{
				Average:             -5,
				AveragePreviousFile: 0,
			},
			args: args{
				data: [][]string{
					{"00:01:02", "-5"},
					{"00:01:05", "-7"},
					{"00:01:07", "-8"},
					{"00:01:08", "-12"},
					{"00:01:12", "-7"},
				},
				baseTime: testFuncDataTime("00:00:00"),
			},
			want:  nil,
			want2: 0,
		},
		{
			name: "NoData",
			fields: fields{
				Average:             0,
				AveragePreviousFile: -4,
			},
			args: args{
				data:     [][]string{},
				baseTime: testFuncDataTime("00:00:00"),
			},
			want:  nil,
			want2: 0,
		},
		{
			name: "usually data",
			fields: fields{
				Average:             0,
				AveragePreviousFile: 0,
			},
			args: args{
				data: [][]string{
					{"00:01:02", "-5"},
					{"00:01:05", "-7"},
					{"00:01:07", "-8"},
					{"00:01:08", "-12"},
					{"00:01:12", "-7"},
				},
				baseTime: testFuncDataTime("00:01:00"),
			},
			want:  []string{"00:01:00", "-8"},
			want2: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Formatter{
				Average:             tt.fields.Average,
				AveragePreviousFile: tt.fields.AveragePreviousFile,
			}
			got, got2 := f.GetAve(tt.args.data, tt.args.baseTime)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAve() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetAve() got2 = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}

func testFuncDataTime(date string) *time.Time {
	t, _ := time.Parse("15:04:05", date)
	return &t
}

func TestFormatter_GetAveAfter(t *testing.T) {
	type fields struct {
		Average             int
		AveragePreviousFile int
	}
	type args struct {
		baseTime *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
		{
			name: "usually test",
			fields: fields{
				Average:             -54,
				AveragePreviousFile: 0,
			},
			args: args{
				baseTime: testFuncDataTime("00:00:00"),
			},
			want: []string{"00:00:00", "-54"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Formatter{
				Average:             tt.fields.Average,
				AveragePreviousFile: tt.fields.AveragePreviousFile,
			}
			if got := f.GetAveAfter(tt.args.baseTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAveAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_GetAvePreviousFile(t *testing.T) {
	type fields struct {
		Average             int
		AveragePreviousFile int
	}
	type args struct {
		baseTime *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
		{
			name: "usually test",
			fields: fields{
				Average:             0,
				AveragePreviousFile: -34,
			},
			args: args{
				baseTime: testFuncDataTime("00:00:00"),
			},
			want: []string{"00:00:00", "-34"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Formatter{
				Average:             tt.fields.Average,
				AveragePreviousFile: tt.fields.AveragePreviousFile,
			}
			if got := f.GetAvePreviousFile(tt.args.baseTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvePreviousFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestFormatter_FormatData(t *testing.T) {
	type fields struct {
		Average             int
		AveragePreviousFile int
	}
	type args struct {
		data [][]string
	}

	f, err := os.Open("test.csv")
	if err != nil {
		t.Errorf("test.csv not found")
	}
	defer f.Close()
	r := csv.NewReader(f)
	alldata, err := r.ReadAll()

	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]string
	}{
		// TODO: Add test cases.
		{
			name:   "test.csv",
			fields: fields{
				Average:             0,
				AveragePreviousFile: 0,
			},
			args:   args{
				data: alldata,
			},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Formatter{
				Average:             tt.fields.Average,
				AveragePreviousFile: tt.fields.AveragePreviousFile,
			}
			if got := f.FormatData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatData() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
