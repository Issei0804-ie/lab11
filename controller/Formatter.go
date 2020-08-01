package controller

import "time"

type Formatter struct {
}

func (f *Formatter) FormatData(data []string) []string {
	panic("not implement")
}

func (f *Formatter) GetAve() *time.Time {
	panic("not implement")
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
