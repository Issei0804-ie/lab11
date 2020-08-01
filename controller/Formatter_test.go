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
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
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
			want: []string{"00:00:00", "-6"},
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
			want:    []string{"00:00:00", "-8"},
			wantErr: false,
		},
		{
			name:    "NoAveData",
			fields:  fields{
				Average: -5,
				AveragePreviousFile: 0,
			},
			args:    args{
				data:    [][]string{
			{"00:01:02", "-5"},
			{"00:01:05", "-7"},
			{"00:01:07", "-8"},
			{"00:01:08", "-12"},
			{"00:01:12", "-7"},
			},
				baseTime: testFuncDataTime("00:00:00"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "NoData",
			fields:  fields{
				Average:             0,
				AveragePreviousFile: -4,
			},
			args:    args{
				data:     [][]string{},
				baseTime: testFuncDataTime("00:00:00"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "usually data",
			fields: fields{
				Average:             0,
				AveragePreviousFile: 0,
			},
			args: args{
				data:    [][]string{
					{"00:01:02", "-5"},
					{"00:01:05", "-7"},
					{"00:01:07", "-8"},
					{"00:01:08", "-12"},
					{"00:01:12", "-7"},
				},
				baseTime: testFuncDataTime("00:01:00"),
			},
			want:    []string{"00:01:00", "-8"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Formatter{
				Average:             tt.fields.Average,
				AveragePreviousFile: tt.fields.AveragePreviousFile,
			}
			got, err := f.GetAve(tt.args.data, tt.args.baseTime)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAve() got = %v, want %v", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func testFuncDataTime(date string) *time.Time {
	t, _ := time.Parse("15:04:05", date)
	return &t
}
