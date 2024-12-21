package geezdate

import (
	"fmt"

	"strings"
	"time"
)

func ExampleConvert() {
	var t time.Time
	t = t.AddDate(2023, 00, 06) //Jan 7 2024
	l := strings.Split(t.String(), " ")
	s := Convert(l[0])
	fmt.Printf("xmass %s\n", s)
	//output: xmass 28-4-2016

}

func ExampleToday() {
	i := Today()
	fmt.Println(i)
	//output: today date
}
func ExampleGeezday() {
	i := Geezday("2024-01-09")
	fmt.Println(i)
	//output: ፴ ታኅሣሥ ፳፻፲፮
}
