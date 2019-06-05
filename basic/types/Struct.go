package main

import(
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat float64
	Lon float64
	Date time.Time
}

func main()  {
	var bootcamp = Bootcamp{Lat: 1, Lon: 2}
	x := new(int)
	fmt.Println(x)
	fmt.Println("Struct")
	fmt.Println(bootcamp)
}
