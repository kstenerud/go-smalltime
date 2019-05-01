package smalltime

import "testing"
import "time"
import "fmt"

func workaroundUnusedImportErrorFmt() {
	fmt.Printf("")
}

func assertEncodeDecode(t *testing.T, year, month, day, hour, minute, second, nsec int) {
	time := NanoNew(year, month, day, hour, minute, second, usec)
	if time.Year() != year || time.Month() != month || time.Day() != day ||
		time.Minute() != minute || time.Second() != second || time.Nanosecond() != nsec {
		t.Errorf("Expected: %04d-%02d-%02dT%02d:%02d:%02d.%06d, Actual: %04d-%02d-%02dT%02d:%02d:%02d.%06d (doy %d)",
			year, month, day, hour, minute, second, nsec,
			time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second(), time.Nanosecond(), time.Doy())
	}
}

func assertEncode(t *testing.T, year, month, day, hour, minute, second, nsec int, encoded NanoSmalltime) {
	time := NanoNew(year, month, day, hour, minute, second, nsec)
	if time != encoded {
		t.Errorf("Expected %04d-%02d-%02dT%02d:%02d:%02d.%06d to encode to %016x. Actual: %016x",
			year, month, day, hour, minute, second, nsec, encoded, time)
	}
}

func assertDecode(t *testing.T, time NanoSmalltime, year, month, day, hour, minute, second, nsec int) {
	if time.Year() != year || time.Month() != month || time.Day() != day ||
		time.Minute() != minute || time.Second() != second || time.Nanosecond() != nsec {
		t.Errorf("Expected %016x to decode to %04d-%02d-%02dT%02d:%02d:%02d.%06d. Actual: %04d-%02d-%02dT%02d:%02d:%02d.%06d",
			time, year, month, day, hour, minute, second, nsec,
			time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second(), time.Nanosecond())
	}
}

func assertYdToMd(t *testing.T, year, doy, month, day int) {
	mm, dd := doyToYmd(year, doy)
	if mm != month || dd != day {
		t.Errorf("Expected Y %d DOY %d to give M %d D %d, but got %d %d", year, doy, month, day, mm, dd)
	}
}

func assertYmdToDoy(t *testing.T, year, month, day, doy int) {
	dd := ymdToDoy(year, month, day)
	if dd != doy {
		t.Errorf("Expected Y %d M %d D %d to give DOY %d, but got %d", year, month, day, doy, dd)
	}
}

func assertTimeEquivalence(t *testing.T, year, month, day, hour, minute, second, nsec int) {
	smtime := NanoNew(year, month, day, hour, minute, second, nsec)
	gotime := time.Date(year, time.Month(month), day, hour, minute, second, nsec, time.UTC)

	if smtime.Year() != gotime.Year() ||
		smtime.Month() != int(gotime.Month()) ||
		smtime.Day() != gotime.Day() ||
		smtime.Hour() != gotime.Hour() ||
		smtime.Minute() != gotime.Minute() ||
		smtime.Second() != gotime.Second() ||
		smtime.Nanosecond() != gotime.Nanosecond() {
		t.Errorf("%04d-%02d-%02dT%02d:%02d:%02d.%06d != %v",
			smtime.Year(), smtime.Month(), smtime.Day(), smtime.Hour(),
			smtime.Minute(), smtime.Second(), smtime.Nanosecond(), gotime)
	}

	if smtime.AsTime() != gotime {
		t.Errorf("%04d-%02d-%02dT%02d:%02d:%02d.%06d did not convert cleanly to %v",
			smtime.Year(), smtime.Month(), smtime.Day(), smtime.Hour(),
			smtime.Minute(), smtime.Second(), smtime.Nanosecond(), gotime)
	}

	if FromTime(gotime) != smtime {
		t.Errorf("%v did not convert cleanly to %04d-%02d-%02dT%02d:%02d:%02d.%06d",
			gotime, smtime.Year(), smtime.Month(), smtime.Day(), smtime.Hour(),
			smtime.Minute(), smtime.Second(), smtime.Nanosecond())
	}
}

func assertGreater(t *testing.T, greater, smaller NanoSmalltime) {
	if greater <= smaller {
		t.Errorf("%04d-%02d-%02dT%02d:%02d:%02d.%06d is not smaller than %04d-%02d-%02dT%02d:%02d:%02d.%06d",
			greater.Year(), greater.Month(), greater.Day(), greater.Hour(),
			greater.Minute(), greater.Second(), greater.Nanosecond(),
			smaller.Year(), smaller.Month(), smaller.Day(), smaller.Hour(),
			smaller.Minute(), smaller.Second(), smaller.Nanosecond())
	}
}

var daysInMonthNormal = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var daysInMonthLeap = [...]int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func testIsLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	return year%4 == 0
}

// ==================================================================

func Test409Years(t *testing.T) {
	for year := 1999; year <= 2408; year++ {
		daysInMonthTable := &daysInMonthNormal
		if testIsLeapYear(year) {
			daysInMonthTable = &daysInMonthLeap
		}
		for month := 1; month <= 12; month++ {
			daysInMonth := daysInMonthTable[month]
			for day := 1; day <= daysInMonth; day++ {
				assertEncodeDecode(t, year, month, day, 0, 0, 0, 0)
			}
		}
	}
}

func TestNanoseconds(t *testing.T) {
	year := 2003
	month := 11
	day := 15
	hour := 8
	minute := 30
	second := 55
	for nanosecond := 0; nanosecond < 1000000000; nanosecond++ {
		assertTimeEquivalence(t, year, month, day, hour, minute, second, nanosecond)
	}
}

func TestSeconds(t *testing.T) {
	year := 2003
	month := 11
	day := 15
	hour := 8
	minute := 30
	nanosecond := 1402778
	for second := 0; second < 60; second++ {
		assertTimeEquivalence(t, year, month, day, hour, minute, second, nanosecond)
	}
}

func TestMinutes(t *testing.T) {
	year := 2003
	month := 11
	day := 15
	hour := 8
	second := 30
	nanosecond := 1402991
	for minute := 0; minute < 60; minute++ {
		assertTimeEquivalence(t, year, month, day, hour, minute, second, nanosecond)
	}
}

func TestHours(t *testing.T) {
	year := 2003
	month := 11
	day := 15
	minute := 30
	second := 55
	nanosecond := 1402889
	for hour := 0; hour < 24; hour++ {
		assertTimeEquivalence(t, year, month, day, hour, minute, second, nanosecond)
	}
}

func TestDays(t *testing.T) {
	year := 2003
	month := 11
	hour := 8
	minute := 30
	second := 55
	nanosecond := 1402345
	for day := 1; day < 30; day++ {
		assertTimeEquivalence(t, year, month, day, hour, minute, second, nanosecond)
	}
	year = 2004
	for day := 1; day < 30; day++ {
		assertTimeEquivalence(t, year, month, day, hour, minute, second, nanosecond)
	}
}

func TestMonths(t *testing.T) {
	year := 2003
	day := 15
	hour := 8
	minute := 30
	second := 55
	nanosecond := 1402988
	for month := 1; month < 12; month++ {
		assertTimeEquivalence(t, year, month, day, hour, minute, second, nanosecond)
	}
	year = 2004
	for month := 1; month < 12; month++ {
		assertTimeEquivalence(t, year, month, day, hour, minute, second, nanosecond)
	}
}

func TestYears(t *testing.T) {
	month := 11
	day := 15
	hour := 8
	minute := 15
	second := 30
	nanosecond := 1402567

	for year := -131072; year < 131071; year++ {
		assertTimeEquivalence(t, year, month, day, hour, minute, second, nanosecond)
	}
}

func TestYmdToDoy(t *testing.T) {
	assertYmdToDoy(t, 1985, 1, 1, 1)
	assertYmdToDoy(t, 1985, 1, 2, 2)
	assertYmdToDoy(t, 1985, 1, 31, 31)
	assertYmdToDoy(t, 1985, 2, 1, 32)
	assertYmdToDoy(t, 1985, 2, 28, 59)
	assertYmdToDoy(t, 1985, 3, 1, 60)
	assertYmdToDoy(t, 1985, 3, 2, 61)
	assertYmdToDoy(t, 1985, 3, 31, 90)
	assertYmdToDoy(t, 1985, 4, 1, 91)
	assertYmdToDoy(t, 1985, 12, 30, 364)
	assertYmdToDoy(t, 1985, 12, 31, 365)
	assertYmdToDoy(t, 1986, 1, 1, 1)

	assertYmdToDoy(t, 2000, 1, 1, 1)
	assertYmdToDoy(t, 2000, 1, 2, 2)
	assertYmdToDoy(t, 2000, 1, 31, 31)
	assertYmdToDoy(t, 2000, 2, 1, 32)
	assertYmdToDoy(t, 2000, 2, 28, 59)
	assertYmdToDoy(t, 2000, 2, 29, 60)
	assertYmdToDoy(t, 2000, 3, 1, 61)
	assertYmdToDoy(t, 2000, 3, 2, 62)
	assertYmdToDoy(t, 2000, 3, 31, 91)
	assertYmdToDoy(t, 2000, 4, 1, 92)
	assertYmdToDoy(t, 2000, 12, 30, 365)
	assertYmdToDoy(t, 2000, 12, 31, 366)
	assertYmdToDoy(t, 2001, 1, 1, 1)

	assertYmdToDoy(t, 2100, 1, 1, 1)
	assertYmdToDoy(t, 2100, 1, 2, 2)
	assertYmdToDoy(t, 2100, 1, 31, 31)
	assertYmdToDoy(t, 2100, 2, 1, 32)
	assertYmdToDoy(t, 2100, 2, 28, 59)
	assertYmdToDoy(t, 2100, 3, 1, 60)
	assertYmdToDoy(t, 2100, 3, 2, 61)
	assertYmdToDoy(t, 2100, 3, 31, 90)
	assertYmdToDoy(t, 2100, 4, 1, 91)
	assertYmdToDoy(t, 2100, 12, 30, 364)
	assertYmdToDoy(t, 2100, 12, 31, 365)
	assertYmdToDoy(t, 2101, 1, 1, 1)
}

func TestDoyToYmd(t *testing.T) {
	assertYdToMd(t, 1985, 1, 1, 1)
	assertYdToMd(t, 1985, 2, 1, 2)
	assertYdToMd(t, 1985, 31, 1, 31)
	assertYdToMd(t, 1985, 32, 2, 1)
	assertYdToMd(t, 1985, 59, 2, 28)
	assertYdToMd(t, 1985, 60, 3, 1)
	assertYdToMd(t, 1985, 61, 3, 2)
	assertYdToMd(t, 1985, 90, 3, 31)
	assertYdToMd(t, 1985, 91, 4, 1)
	assertYdToMd(t, 1985, 364, 12, 30)
	assertYdToMd(t, 1985, 365, 12, 31)
	assertYdToMd(t, 1986, 1, 1, 1)

	assertYdToMd(t, 2000, 1, 1, 1)
	assertYdToMd(t, 2000, 2, 1, 2)
	assertYdToMd(t, 2000, 31, 1, 31)
	assertYdToMd(t, 2000, 32, 2, 1)
	assertYdToMd(t, 2000, 59, 2, 28)
	assertYdToMd(t, 2000, 60, 2, 29)
	assertYdToMd(t, 2000, 61, 3, 1)
	assertYdToMd(t, 2000, 62, 3, 2)
	assertYdToMd(t, 2000, 91, 3, 31)
	assertYdToMd(t, 2000, 92, 4, 1)
	assertYdToMd(t, 2000, 365, 12, 30)
	assertYdToMd(t, 2000, 366, 12, 31)
	assertYdToMd(t, 2001, 1, 1, 1)

	assertYdToMd(t, 2100, 1, 1, 1)
	assertYdToMd(t, 2100, 2, 1, 2)
	assertYdToMd(t, 2100, 31, 1, 31)
	assertYdToMd(t, 2100, 32, 2, 1)
	assertYdToMd(t, 2100, 59, 2, 28)
	assertYdToMd(t, 2100, 60, 3, 1)
	assertYdToMd(t, 2100, 61, 3, 2)
	assertYdToMd(t, 2100, 90, 3, 31)
	assertYdToMd(t, 2100, 91, 4, 1)
	assertYdToMd(t, 2100, 364, 12, 30)
	assertYdToMd(t, 2100, 365, 12, 31)
	assertYdToMd(t, 2101, 1, 1, 1)
}

func TestComparisons(t *testing.T) {
	assertGreater(t, NanoNew(2000, 1, 1, 0, 0, 0, 1), NanoNew(2000, 1, 1, 0, 0, 0, 0))
	assertGreater(t, NanoNew(2000, 1, 1, 0, 0, 1, 0), NanoNew(2000, 1, 1, 0, 0, 0, 999999))
	assertGreater(t, NanoNew(2000, 1, 1, 0, 0, 2, 0), NanoNew(2000, 1, 1, 0, 0, 1, 0))
	assertGreater(t, NanoNew(2000, 1, 1, 0, 1, 0, 0), NanoNew(2000, 1, 1, 0, 0, 60, 0))
	assertGreater(t, NanoNew(2000, 1, 1, 0, 2, 0, 0), NanoNew(2000, 1, 1, 0, 1, 0, 0))
	assertGreater(t, NanoNew(2000, 1, 1, 1, 0, 0, 0), NanoNew(2000, 1, 1, 0, 59, 0, 0))
	assertGreater(t, NanoNew(2000, 1, 1, 2, 0, 0, 0), NanoNew(2000, 1, 1, 1, 0, 0, 0))
	assertGreater(t, NanoNew(2000, 1, 2, 0, 0, 0, 0), NanoNew(2000, 1, 1, 23, 0, 0, 0))
	assertGreater(t, NanoNew(2000, 1, 2, 0, 0, 0, 0), NanoNew(2000, 1, 1, 0, 0, 0, 0))
	assertGreater(t, NanoNew(2005, 1, 1, 0, 0, 0, 0), NanoNew(2004, 12, 31, 0, 0, 0, 0))
	assertGreater(t, NanoNew(1, 1, 1, 0, 0, 0, 0), NanoNew(0, 1, 1, 0, 0, 0, 0))
	assertGreater(t, NanoNew(0, 1, 1, 0, 0, 0, 0), NanoNew(-1, 1, 1, 0, 0, 0, 0))
	assertGreater(t, NanoNew(1, 1, 1, 0, 0, 0, 0), NanoNew(-1, 1, 1, 0, 0, 0, 0))
}

func TestSpecExamples(t *testing.T) {
	assertDecode(t, NanoSmalltime(0x1f06b48590dbc2e), 1985, 10, 26, 8, 22, 16, 900142)
	assertDecode(t, NanoSmalltime(0x1f06b68590dbc2e), 1985, 10, 27, 8, 22, 16, 900142)
	assertDecode(t, NanoSmalltime(0x1f06b48550dbc2e), 1985, 10, 26, 8, 21, 16, 900142)

	assertEncode(t, 1985, 10, 26, 8, 22, 16, 900142, NanoSmalltime(0x1f06b48590dbc2e))
	assertEncode(t, 1985, 10, 27, 8, 22, 16, 900142, NanoSmalltime(0x1f06b68590dbc2e))
	assertEncode(t, 1985, 10, 26, 8, 21, 16, 900142, NanoSmalltime(0x1f06b48550dbc2e))
}
