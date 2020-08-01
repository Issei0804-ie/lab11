package controller

import (
	"fmt"
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
	panic("not imprement")
}

func (f *Formatter) GetAve(data [][]string, baseTime *time.Time) ([]string, error) {
	ave := 0
	i := 0
	subTime := baseTime.Add(time.Second*10)
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
	if ave == 0{
		return nil, fmt.Errorf("Average cannot be defined")
	}
	records := []string{f.timeToData(*baseTime), strconv.Itoa(ave / i)}
	return records, nil
}

func (f *Formatter) GetAveAfter() *time.Time {
	panic("not implement")
}

func (f *Formatter) GetAvePreviousFile() *time.Time {
	panic("not implement")
}

func (f *Formatter) SetAve(time *time.Time) {
	panic("not implement")
}

func (f *Formatter) SetAvePreviousFile(time *time.Time) {
	panic("not implement")
}

func (f *Formatter) dataToTime(date string) time.Time {
	t, _ := time.Parse("15:04:05", date)
	return t
}

func (f *Formatter) timeToData(t time.Time) string {
	return t.Format("15:04:05")
}

func (f *Formatter) timeToSeconds(t time.Time) int{
	return (t.Hour()*60*60) + (t.Minute()*60) + t.Second()
}
