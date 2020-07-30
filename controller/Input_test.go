package controller

import (
	"encoding/csv"
	"os"
	"reflect"
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
			dc := &DirConf{
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

func TestInput_ReadCSV(t *testing.T) {
	type fields struct {
		data [][]string
		conf config
	}
	type args struct {
		filepath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]string
	}{
		// TODO: Add test cases.
		{
			"ReadCSVTEST",
			fields{
				data: nil,
				conf: config{},
			},
			args{
				filepath: "removeHeadTestData.csv",
			},
			[][]string{
				{"Body", "Hoge"},
				{"Body2", "Fuga"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := &Input{
				data: tt.fields.data,
				conf: tt.fields.conf,
			}
			if got := in.ReadCSV(tt.args.filepath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInput_GetYearMonth(t *testing.T) {
	type fields struct {
		data [][]string
		conf config
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
		{
			"GetYearMonth",
			fields{
				data: nil,
				conf: newConfig(),
			},
			[]string{"200906", "200910", "200911", "200912"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := &Input{
				data: tt.fields.data,
				conf: tt.fields.conf,
			}
			if got := in.GetYearMonth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetYearMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInput_GetDirNames(t *testing.T) {
	type fields struct {
		data [][]string
		conf config
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]string
	}{
		// TODO: Add test cases.
		{
			name: "DirNames",
			fields: fields{
				data: nil,
				conf: newConfig(),
			},
			want: [][]string{
				{"20090601", "20090602", "20090603", "20090604", "20090605", "20090606", "20090607", "20090608", "20090609", "20090610", "20090611", "20090612", "20090613", "20090614", "20090615", "20090616", "20090617", "20090618", "20090619", "20090620", "20090621", "20090622", "20090623", "20090624", "20090625", "20090626", "20090627", "20090628", "20090629", "20090630"},
				{"20091001", "20091002", "20091003", "20091004", "20091005", "20091006", "20091007", "20091008", "20091009", "20091010", "20091011", "20091012", "20091013", "20091014", "20091015", "20091016", "20091017", "20091018", "20091019", "20091020", "20091021", "20091022", "20091023", "20091024", "20091025", "20091026", "20091027", "20091028", "20091029", "20091030", "20091031"},
				{"20091101", "20091102", "20091103", "20091104", "20091105", "20091106", "20091107", "20091108", "20091109", "20091110", "20091111", "20091112", "20091113", "20091114", "20091115", "20091116", "20091117", "20091118", "20091119", "20091120", "20091121", "20091122", "20091123", "20091124", "20091125", "20091126", "20091127", "20091128", "20091129", "20091130"},
				{"20091201", "20091202", "20091203", "20091204", "20091205", "20091206", "20091207", "20091208", "20091209", "20091210", "20091211", "20091212", "20091213", "20091214", "20091215", "20091216", "20091217", "20091218", "20091219", "20091220", "20091221", "20091222", "20091223", "20091224", "20091225", "20091226", "20091227", "20091228", "20091229", "20091230", "20091231"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := &Input{
				data: tt.fields.data,
				conf: tt.fields.conf,
			}
			if got := in.GetDirNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDirNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
