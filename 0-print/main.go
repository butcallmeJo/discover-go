package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Hello, we are Holberton School")
	fmt.Println("the date is ", t)
	fmt.Println("the year is ", t.Year())
	fmt.Println("the month is", t.Month())
	fmt.Println("the day is", t.Day())
}
