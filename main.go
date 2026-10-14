package main

// type Data struct {
// 	Min   float64
// 	Max   float64
// 	Sum   float64
// 	Count int
// }

func main() {
	// start := time.Now()

	// processConcurrent("weather_stations.csv")
	prosessChunking("weather_stations.csv")

	// stationData := make(map[string]*Data)
	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	line := strings.TrimSpace(scanner.Text())
	// 	parts := strings.Split(line, ";")

	// 	town := strings.TrimSpace(parts[0])
	// 	temp, err := strconv.ParseFloat(parts[1], 64)

	// 	if err != nil {
	// 		fmt.Println("Error Convert Float")
	// 	}

	// 	stats, exist := stationData[town]
	// 	if !exist {
	// 		stationData[town] = &Data{
	// 			Min:   temp,
	// 			Max:   temp,
	// 			Sum:   temp,
	// 			Count: 1,
	// 		}
	// 	} else {
	// 		if temp < stats.Min {
	// 			stats.Min = temp
	// 		}
	// 		if temp > stats.Max {
	// 			stats.Max = temp
	// 		}
	// 		stats.Sum += temp
	// 		stats.Count++
	// 	}
	// }

	// for city, stats := range stationData {
	// 	avg := stats.Sum / float64(stats.Count)
	// 	fmt.Printf("Kota : %s, Avarage : %.4f, Min : %.4f, Max : %.4f\n", city, avg, stats.Min, stats.Max)
	// }

	// elapsed := time.Since(start)
	// fmt.Printf("Waktu : %v", elapsed)

}
