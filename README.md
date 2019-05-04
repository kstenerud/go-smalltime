Smalltime & Nanotime
====================

Go implementations of [smalltime](https://github.com/kstenerud/smalltime/blob/master/smalltime-specification.md)
and [nanotime](https://github.com/kstenerud/smalltime/blob/master/nanotime-specification.md).

Smalltime & Nanotime are a simple and convenient binary date & time formats in 64 bits.


Smalltime Features
------------------

 * Encodes a complete date & time into a 64-bit integer.
 * Fields are compatible with ISO-8601.
 * Maintenance-free (no leap second tables to update).
 * Easily converts to human readable fields.
 * Supports hundreds of thousands of years (Smalltime).
 * Supports time units to the microsecond (Smalltime) or nanosecond (Nanotime).
 * Supports leap years and leap seconds.
 * Encoded values are comparable (within the same type).

Smalltime is based off a signed 64-bit integer. It has a range of hundreds of
thousands of years, but only goes down to the microsecond.

Nanotime is based off an unsigned 64-bit integer. It goes down to the
nanosecond, but only has a range of 256 years (1970 - 2226).


Library Usage
-------------

```golang
import "fmt"
import "github.com/kstenerud/go-smalltime"

func demonstrateSmalltime() {
	noonJan12000 := smalltime.NewSmalltime(2000, 1, 1, 12, 0, 0, 0)
	oneOclockJan12000 := smalltime.NewSmalltime(2000, 1, 1, 13, 0, 0, 0)
	feb151999 := smalltime.NewSmalltime(1999, 2, 15, 12, 8, 45, 9122)

	if oneOclockJan12000 > noonJan12000 {
		fmt.Printf("Comparison: One o'clock is greater than noon.\n")
	}

	if feb151999 < noonJan12000 {
		fmt.Printf("Comparison: Feb 15, 1999 is less than Jan 1, 2000.\n")
	}

	gotime := feb151999.AsTime()
	fmt.Printf("Go Time: %v\n", gotime)

	smtime := smalltime.SmalltimeFromTime(gotime)
	fmt.Printf("Smalltime Raw: 0x%016x\n", smtime)

	fmt.Printf("Smalltime Fields: %04d-%02d-%02d %02d:%02d:%02d.%06d\n",
		smtime.Year(), smtime.Month(), smtime.Day(), smtime.Hour(),
		smtime.Minute(), smtime.Second(), smtime.Microsecond())
}
```

Output:

```
Comparison: One o'clock is greater than noon.
Comparison: Feb 15, 1999 is less than Jan 1, 2000.
Go Time: 1999-02-15 12:08:45.009122 +0000 UTC
Smalltime Raw: 0x01f3c9ec22d023a2
Smalltime Fields: 1999-02-15 12:08:45.009122
```

```golang
import "fmt"
import "github.com/kstenerud/go-smalltime"

func demonstrateNanotime() {
	noonJan12000 := smalltime.NewNanotime(2000, 1, 1, 12, 0, 0, 0)
	oneOclockJan12000 := smalltime.NewNanotime(2000, 1, 1, 13, 0, 0, 0)
	feb151999 := smalltime.NewNanotime(1999, 2, 15, 12, 8, 45, 10159122)

	if oneOclockJan12000 > noonJan12000 {
		fmt.Printf("Comparison: One o'clock is greater than noon.\n")
	}

	if feb151999 < noonJan12000 {
		fmt.Printf("Comparison: Feb 15, 1999 is less than Jan 1, 2000.\n")
	}

	gotime := feb151999.AsTime()
	fmt.Printf("Go Time: %v\n", gotime)

	smtime := smalltime.NanotimeFromTime(gotime)
	fmt.Printf("Nanotime Raw: 0x%016x\n", smtime)

	fmt.Printf("Nanotime Fields: %04d-%02d-%02d %02d:%02d:%02d.%09d\n",
		smtime.Year(), smtime.Month(), smtime.Day(), smtime.Hour(),
		smtime.Minute(), smtime.Second(), smtime.Nanosecond())
}
```

Output:

```
Comparison: One o'clock is greater than noon.
Comparison: Feb 15, 1999 is less than Jan 1, 2000.
Go Time: 1999-02-15 12:08:45.010159122 +0000 UTC
Nanotime Raw: 0x1d27b08b409b0412
Nanotime Fields: 1999-02-15 12:08:45.010159122
```
