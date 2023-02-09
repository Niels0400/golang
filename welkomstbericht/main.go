package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	hour := currentTime.Hour()

	var groete string
	if hour >= 7 && hour < 12 {
		groete = "Goedemorgen"
	} else if hour >= 12 && hour < 18 {
		groete = "Goedemiddag"
	} else if hour >= 18 && hour < 23 {
		groete = "Goedenavond"
	} else {
		groete = "Sorry, de parkeerplaats is 's nachts gesloten"
	}
	fmt.Println(groete + "! Welkom bij Fonteyn Vakantieparken")
}
