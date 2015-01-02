package main

import "time"

var months = map[time.Month]string{
	time.January: 	"01 - Январь",
	time.February: 	"02 - Февраль",
	time.March:		"03 - Март",
    time.April:		"04 - Апрель",
    time.May:		"05 - Май",
    time.June:		"06 - Июнь",
    time.July:		"07 - Июль",
    time.August:	"08 - Август",
    time.September:	"09 - Сентябрь",
    time.October:	"10 - Октябрь",
    time.November:	"11 - Ноябрь",
    time.December:	"12 - Декабрь",
}

func GetMonthsName(month time.Month) string {
	return months[month]
}
