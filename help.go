package smalltime
//import "fmt"

func WrapNanoNew(year, month, day, hour, minute, second, nanosecond int) NanoSmalltime {
//fmt.Println("A",year, month, day, hour, minute, second, nanosecond)
    n:=NanoNew(year, month, day, hour, minute, second, nanosecond) 
//fmt.Println("B",n.Year(),n.Month(),n.Day())
return n
}
