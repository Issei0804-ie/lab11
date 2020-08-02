package controller

import (
	"strconv"
	"time"
)

const (
	TIME    = 0
	RXLEVEL = 1
)

type Formatter struct {
	Average             int
	AveragePreviousFile int
}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) FormatData(data [][]string) [][]string {
	baseTime := f.dataToTime("00:00:00")
	sliceCount := 0
	var records [][]string
	for {
		record, index := f.GetAve(data[sliceCount:], &baseTime)
		if record == nil {
			record = f.GetAveAfter(&baseTime)
			if record == nil {
				record = f.GetAvePreviousFile(&baseTime)
			}
		} else {
			sliceCount += index
		}
		records = append(records, record)
		if baseTime == f.dataToTime("23:50:50") {
			break
		}
		baseTime = baseTime.Add(time.Second * 10)
	}
	tmp := 0
	i := 0
	for i = 0; i < len(records)-1; i++ {
		work, _ := strconv.Atoi(records[i][RXLEVEL])
		tmp += work
	}
	f.SetAvePreviousFile(tmp / i)
	return records
}

func (f *Formatter) GetAve(data [][]string, baseTime *time.Time) ([]string, int) {
	ave := 0
	i := 0
	subTime := baseTime.Add(time.Second * 10)
	for i = 0; i < len(data); i++ {
		tmp := f.timeToSeconds(subTime) - f.timeToSeconds(f.dataToTime(data[i][TIME]))
		if 1 <= tmp && tmp <= 10 {
			n, err := strconv.Atoi(data[i][RXLEVEL])
			if err != nil {
				panic(err.Error())
			}
			ave += n
		} else {
			break
		}
	}
	if ave == 0 {
		return nil, 0
	}
	records := []string{f.timeToData(*baseTime), strconv.Itoa(ave / i)}
	f.Average = ave / i
	return records, i
}

func (f *Formatter) GetAveAfter(baseTime *time.Time) []string {
	return []string{f.timeToData(*baseTime), strconv.Itoa(f.Average)}
}

func (f *Formatter) GetAvePreviousFile(baseTime *time.Time) []string {
	return []string{f.timeToData(*baseTime), strconv.Itoa(f.AveragePreviousFile)}
}

func (f *Formatter) SetAve(ave int) {
	f.Average = ave
}

func (f *Formatter) SetAvePreviousFile(ave int) {
	f.AveragePreviousFile = ave
}

func (f *Formatter) dataToTime(date string) time.Time {
	t, _ := time.Parse("15:04:05", date)
	return t
}

func (f *Formatter) timeToData(t time.Time) string {
	return t.Format("15:04:05")
}

func (f *Formatter) timeToSeconds(t time.Time) int {
	return (t.Hour() * 60 * 60) + (t.Minute() * 60) + t.Second()
}
