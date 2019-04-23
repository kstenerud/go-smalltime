/*
Smalltime is a binary date & time representation that fits into a 64-bit signed integer.
Smalltime values are directly comparable, and the standard Gregorian fields are easily extractable.

Specification: https://github.com/kstenerud/smalltime/blob/master/smalltime-specification.md


Features
--------

 * Encodes a complete date & time into a 64-bit signed integer.
 * Fields (including year) are compatible with ISO-8601.
 * Maintenance-free (no leap second tables to update).
 * Easily converts to human readable fields.
 * Supports hundreds of thousands of years.
 * Supports time units to the microsecond.
 * Supports leap years and leap seconds.
 * Encoded values are comparable.
*/
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

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func ymdToDoy(year, month, day int) (doy int) {
	monthsFromMarch := (month + 9) % 12       // [0, 11]
	doy = (153*monthsFromMarch+2)/5 + day - 1 // [0, 365]
	if isLeapYear(year) {
		doy = (doy + 60) % 366
	} else {
		doy = (doy + 59) % 365
	}
	doy += 1
	return doy // [1, 366]
}

func doyToYmd(year, doy int) (month, day int) {
	if isLeapYear(year) {
		doy = (doy + 305) % 366
	} else {
		doy = (doy + 305) % 365
	}
	monthsFromMarch := (5*doy + 2) / 153      // [0, 11]
	day = doy - (153*monthsFromMarch+2)/5 + 1 // [1, 31]
	month = (monthsFromMarch+2)%12 + 1        // [1, 12]
	return month, day
}

func FromTime(t time.Time) Smalltime {
	return New(t.Year(), int(t.Month()), t.Day(), t.Hour(),
		t.Minute(), t.Second(), t.Nanosecond()/1000)
}

func New(year, month, day, hour, minute, second, microsecond int) Smalltime {
	return Smalltime(year)<<bitshiftYear |
		Smalltime(month)<<bitshiftMonth |
		Smalltime(day)<<bitshiftDay |
		Smalltime(hour)<<bitshiftHour |
		Smalltime(minute)<<bitshiftMinute |
		Smalltime(second)<<bitshiftSecond |
		Smalltime(microsecond)
}

func NewWithDoy(year, dayOfYear, hour, minute, second, microsecond int) Smalltime {
	month, day := doyToYmd(year, dayOfYear)
	return New(year, month, day, hour, minute, second, microsecond)
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
