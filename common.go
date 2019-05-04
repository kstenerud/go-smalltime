/*
Smalltime and Nanotime are binary date & time representations that fit into
64-bit integers. Values are directly comparable (within the same type), and the
standard Gregorian fields are easily extractable.

Smalltime is based off a signed 64-bit integer. It has a range of hundreds of
thousands of years, but only goes down to the microseconds.

Nanotime is based off an unsigned 64-bit integer. It goes down to the
nanosecond, but only has a range of 256 years (1970 - 2226).


Specifications:

 * https://github.com/kstenerud/smalltime/blob/master/smalltime-specification.md
 * https://github.com/kstenerud/smalltime/blob/master/nanotime-specification.md


Features
--------

 * Encodes a complete date & time into a 64-bit integer.
 * Fields are compatible with ISO-8601.
 * Maintenance-free (no leap second tables to update).
 * Easily converts to human readable fields.
 * Supports hundreds of thousands of years (Smalltime only).
 * Supports time units to the nanosecond (Nanotime only).
 * Supports leap years and leap seconds.
 * Encoded values are comparable.
*/
package smalltime

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
