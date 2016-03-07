package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())
	num := rand.Intn(100)
	sign := "less"
	if num > 50 {
		sign = "greater"
	}
	school := "Holberton School"
	beautifulWeather := true
	holbertonFounders := []string{"Rudy Rigot", "Julien Barbier", "Sylvain Kalache"}
	fmt.Println("my random number is ", num, " and is ", sign, " than 50")
	fmt.Println("I am a student of the ", school)
	if beautifulWeather == true {
		fmt.Println("It's a beautiful weather!")
	}
	for i := 0; i < 3; i++ {
		fmt.Println(holbertonFounders[i], " is a founder")
	}
}
