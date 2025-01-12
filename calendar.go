// Package geezdate implements functions for converting Gregorian calendar to Geez calander.
package geezdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tgiday/mgn2"
)

type Gdate struct {
	d, m, y int
}

// Geezday return a geez date (eg "2024-01-09" to "፴ ታኅሣሥ ፳፻፲፮") taking Gregorian calander date of format yyyy-mm-dd
func Geezday(date string) Gdate {
	s := Convert(date)
	lst := strings.Split(s, "-")
	g := Gdate{}
	g.d, _ = strconv.Atoi(lst[0])
	g.m, _ = strconv.Atoi(lst[1])
	g.y, _ = strconv.Atoi(lst[2])
	return g
}

// Today return a todays date according to Geez calander
func Today() Gdate {
	t := time.Now()
	ls := strings.Split(t.String(), " ")
	td := ls[0]
	s := Convert(td)
	lst := strings.Split(s, "-")
	g := Gdate{}
	g.d, _ = strconv.Atoi(lst[0])
	g.m, _ = strconv.Atoi(lst[1])
	g.y, _ = strconv.Atoi(lst[2])

	return g
}

func (g Gdate) String() string {
	month := []string{"መስከረም", "ጥቅምት", "ኅዳር", "ታኅሣሥ", "ጥር", "የካቲት", "መጋቢት", "ሚያዝያ", "ግንቦት", "ሰኔ", "ሐምሌ", "ነሐሴ", "ጳጉሜ"}
	d := mgn2.Fmtint(g.d)
	m := month[g.m-1]
	y := mgn2.Fmtint(g.y)
	str := fmt.Sprintf("%s %s %s", d, m, y)
	return str
}

// daysBefore[m] counts the number of days in a non-leap year
// before month m begins. There is an entry for m=13, counting
// the number of days before Meskerem of next year (365).
var daysBefore = [...]int32{
	0,
	30,
	30 + 30,
	30 + 30 + 30,
	30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 5,
}

func convert(x time.Time) (int, int, int) {
	yr := x.Year()
	day := x.YearDay() - 1
	month := int(day / 30)
	end := int(daysBefore[month+1])
	var begin int
	if day >= end {
		month++
		begin = end
	} else {
		begin = int(daysBefore[month])
	}
	month++
	day = day - begin + 1
	if month > 13 {
		month = month - 13
		return month, day, yr
	}
	return month, day, yr
}

// Convert return a Geez calander date ,take string Gregorian calendar date  ("1991-05-24" to 16-9-1983)
func Convert(date string) string {
	d, _ := time.Parse("2006-01-02", date)
	x := d.AddDate(0, 0, -2810)
	if isLeap(x.Year()) {
		c := x.AddDate(0, 0, -1)
		mm, dd, yy := convert(c)
		str := fmt.Sprintf("%v-%v-%v", dd, mm, yy)
		return str
	}
	mm, dd, yy := convert(x)
	str := fmt.Sprintf("%v-%v-%v", dd, mm, yy)
	return str
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
