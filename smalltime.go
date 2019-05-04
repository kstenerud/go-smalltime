package smalltime

import "time"

type Smalltime int64

const bitshiftYear = 46
const bitshiftMonth = 42
const bitshiftDay = 37
const bitshiftHour = 32
const bitshiftMinute = 26
const bitshiftSecond = 20

const maskYear = uint64(0x3ffff) << bitshiftYear
const maskMonth = Smalltime(0xf) << bitshiftMonth
const maskDay = Smalltime(0x1f) << bitshiftDay
const maskHour = Smalltime(0x1f) << bitshiftHour
const maskMinute = Smalltime(0x3f) << bitshiftMinute
const maskSecond = Smalltime(0x3f) << bitshiftSecond
const maskMicrosecond = Smalltime(0xfffff)

func SmalltimeFromTime(t time.Time) Smalltime {
	return NewSmalltime(t.Year(), int(t.Month()), t.Day(), t.Hour(),
		t.Minute(), t.Second(), t.Nanosecond()/1000)
}

func NewSmalltime(year, month, day, hour, minute, second, microsecond int) Smalltime {
	return Smalltime(year)<<bitshiftYear |
		Smalltime(month)<<bitshiftMonth |
		Smalltime(day)<<bitshiftDay |
		Smalltime(hour)<<bitshiftHour |
		Smalltime(minute)<<bitshiftMinute |
		Smalltime(second)<<bitshiftSecond |
		Smalltime(microsecond)
}

func NewSmalltimeWithDoy(year, dayOfYear, hour, minute, second, microsecond int) Smalltime {
	month, day := doyToYmd(year, dayOfYear)
	return NewSmalltime(year, month, day, hour, minute, second, microsecond)
}

func (t Smalltime) AsTime() time.Time {
	return t.AsTimeInLocation(time.UTC)
}

func (t Smalltime) AsTimeInLocation(loc *time.Location) time.Time {
	return time.Date(t.Year(), time.Month(t.Month()), t.Day(), t.Hour(),
		t.Minute(), t.Second(), t.Microsecond()*1000, loc)
}

func (time Smalltime) Year() int {
	return int(time >> bitshiftYear)
}

func (time Smalltime) Doy() int {
	return ymdToDoy(time.Year(), time.Month(), time.Day())
}

func (time Smalltime) Month() int {
	return int((time & maskMonth) >> bitshiftMonth)
}

func (time Smalltime) Day() int {
	return int((time & maskDay) >> bitshiftDay)
}

func (time Smalltime) Hour() int {
	return int((time & maskHour) >> bitshiftHour)
}

func (time Smalltime) Minute() int {
	return int((time & maskMinute) >> bitshiftMinute)
}

func (time Smalltime) Second() int {
	return int((time & maskSecond) >> bitshiftSecond)
}

func (time Smalltime) Microsecond() int {
	return int(time & maskMicrosecond)
}
