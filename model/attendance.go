package model

type Attendance struct {
	Date string
	WorkStart string
	WorkEnd string
	BreakStart string
	BreakEnd string
}

func NewAttendance(d string, ws string, we string, bs string, be string) *Attendance{
	return &Attendance{
		Date: d,
		WorkStart: ws,
		WorkEnd: we,
		BreakStart: bs,
		BreakEnd: be,
	}
}
