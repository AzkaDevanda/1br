package main

import (
	"fmt"
	"os"
)

type Data struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func main() {
	file, err := os.Open("weather_stations.csv")
	if err != nil {
		fmt.Println("Error Membaca File")
		return
	}
	defer file.Close()
}
