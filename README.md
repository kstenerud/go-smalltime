Smalltime
=========

A go implementation of [smalltime](https://github.com/kstenerud/smalltime/blob/master/smalltime-specification.md).

Smalltime is a simple and convenient binary date & time format in 64 bits.


Smalltime Features
------------------

 * Encodes a complete date & time into a 64-bit signed integer.
 * Fields (including year) are compatible with ISO-8601.
 * Maintenance-free (no leap second tables to update).
 * Easily converts to human readable fields.
 * Supports hundreds of thousands of years.
 * Supports time units to the microsecond.
 * Supports leap years and leap seconds.
 * Encoded values are comparable.



Library Usage
-------------

```golang
import "fmt"
import "github.com/kstenerud/go-smalltime"

func demonstrate_smalltime() {
    noon_jan1_2000 := smalltime.New(2000, 1, 1, 12, 0, 0, 0)
    one_oclock_jan1_2000 := smalltime.New(2000, 1, 1, 13, 0, 0, 0)
    feb15_1999 := smalltime.New(1999, 2, 15, 12, 8, 45, 9122)

    if one_oclock_jan1_2000 > noon_jan1_2000 {
        fmt.Printf("Comparison: One o'clock is greater than noon.\n")
    }

    if feb15_1999 < noon_jan1_2000 {
        fmt.Printf("Comparison: Feb 15, 1999 is less than Jan 1, 2000.\n")
    }

    gotime := feb15_1999.AsTime()
    fmt.Printf("Go Time: %v\n", gotime)

    smtime := smalltime.FromTime(gotime)
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
