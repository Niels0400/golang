package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	hour := currentTime.Hour()

	var groet string
	if hour >= 7 && hour < 12 {
		groet = "Goedemorgen"
	} else if hour >= 12 && hour < 18 {
		groet = "Goedemiddag"
	} else if hour >= 18 && hour < 23 {
		groet = "Goedenavond"
	} else {
		groet = "Sorry, de parkeerplaats is 's nachts gesloten"
	}
	fmt.Println(groet + "! Welkom bij Fonteyn Vakantieparken")
}
