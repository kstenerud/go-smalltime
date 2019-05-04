package smalltime

import "time"

type Nanotime uint64

const zeroYearNanotime = 1970

const bitshiftYearNanotime = 56
const bitshiftMonthNanotime = 52
const bitshiftDayNanotime = 47
const bitshiftHourNanotime = 42
const bitshiftMinuteNanotime = 36
const bitshiftSecondNanotime = 30

const maskYearNanotime = Nanotime(0xff) << (bitshiftYearNanotime)
const maskMonthNanotime = Nanotime(0xf) << (bitshiftMonthNanotime)
const maskDayNanotime = Nanotime(0x1f) << (bitshiftDayNanotime)
const maskHourNanotime = Nanotime(0x1f) << (bitshiftHourNanotime)
const maskMinuteNanotime = Nanotime(0x3f) << (bitshiftMinuteNanotime)
const maskSecondNanotime = Nanotime(0x3f) << (bitshiftSecondNanotime)
const maskNanoNanotime = Nanotime(0x3fffffff)

func NanotimeFromTime(t time.Time) Nanotime {
	return NewNanotime(t.Year(), int(t.Month()), t.Day(), t.Hour(),
		t.Minute(), t.Second(), t.Nanosecond())
}

func NewNanotime(year, month, day, hour, minute, second, nanosecond int) Nanotime {
	return Nanotime(year-zeroYearNanotime)<<(bitshiftYearNanotime) |
		Nanotime(month)<<(bitshiftMonthNanotime) |
		Nanotime(day)<<(bitshiftDayNanotime) |
		Nanotime(hour)<<(bitshiftHourNanotime) |
		Nanotime(minute)<<(bitshiftMinuteNanotime) |
		Nanotime(second)<<(bitshiftSecondNanotime) |
		Nanotime(nanosecond)
}

func NewNanotimeWithDoy(year, dayOfYear, hour, minute, second, nanosecond int) Nanotime {
	month, day := doyToYmd(year, dayOfYear)
	return NewNanotime(year, month, day, hour, minute, second, nanosecond)
}

func (t Nanotime) AsTime() time.Time {
	return t.AsTimeInLocation(time.UTC)
}

func (t Nanotime) AsTimeInLocation(loc *time.Location) time.Time {
	return time.Date(t.Year(), time.Month(t.Month()), t.Day(), t.Hour(),
		t.Minute(), t.Second(), t.Nanosecond(), loc)
}

func (time Nanotime) Year() int {
	return int(time>>bitshiftYearNanotime) + zeroYearNanotime
}

func (time Nanotime) Doy() int {
	return ymdToDoy(time.Year(), time.Month(), time.Day())
}

func (time Nanotime) Month() int {
	return int((time & maskMonthNanotime) >> (bitshiftMonthNanotime))
}

func (time Nanotime) Day() int {
	return int((time & maskDayNanotime) >> (bitshiftDayNanotime))
}

func (time Nanotime) Hour() int {
	return int((time & maskHourNanotime) >> (bitshiftHourNanotime))
}

func (time Nanotime) Minute() int {
	return int((time & maskMinuteNanotime) >> (bitshiftMinuteNanotime))
}

func (time Nanotime) Second() int {
	return int((time & maskSecondNanotime) >> (bitshiftSecondNanotime))
}

func (time Nanotime) Nanosecond() int {
	return int(time & maskNanoNanotime)
}
