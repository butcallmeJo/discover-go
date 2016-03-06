package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	layout := "January 2, 2006"
	then, _ := time.Parse(layout, u.DOB)
	ageDuration := time.Since(then)
	ageHours := ageDuration.Hours()
	age := int(ageHours / 8760)
	fmt.Println("Hello", u.Name)
	fmt.Println(u.Name, " who was born in ", u.City, " would be ", age, " years old today.")

}
