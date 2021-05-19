package main

import (
    "fmt"
    "github.com/grsmv/goweek"
)

func days() {
    week, _ := goweek.NewWeek(2015, 46)
    fmt.Println(week.Days)
}
