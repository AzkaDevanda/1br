package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Data struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func prosessChunking(filename string) {
	start := time.Now()

	// 1. cek ukuran file
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error Stat File:", err)
		return
	}

	fileSize := fileInfo.Size() // total byte (misal ~18 MB)

	newWorkers := runtime.NumCPU()             // 20 workers
	chunckSize := fileSize / int64(newWorkers) // jatah byte per worker

	var wg sync.WaitGroup
	result := make(chan map[string]*Data, newWorkers)

	for i := 0; i < newWorkers; i++ {
		starts := int64(i) * chunckSize
		end := starts + chunckSize
		if i == newWorkers-1 {
			end = fileSize
		}

		wg.Add(1)
		go processChunk(filename, starts, end, result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	finalData := make(map[string]*Data)

	for localMap := range result {
		for city, localStats := range localMap {
			globalStats, exist := finalData[city]
			if !exist {
				finalData[city] = &Data{
					Min:   localStats.Min,
					Max:   localStats.Max,
					Sum:   localStats.Sum,
					Count: localStats.Count,
				}
			} else {
				if localStats.Min < globalStats.Min {
					globalStats.Min = localStats.Min
				}
				if localStats.Max > globalStats.Max {
					globalStats.Max = localStats.Max
				}
				globalStats.Sum += localStats.Sum
				globalStats.Count += localStats.Count
			}
		}
	}

	elapsed := time.Since(start)

	for city, stats := range finalData {
		avg := stats.Sum / float64(stats.Count)
		fmt.Printf("Kota: %s, Avarage Temp: %.4f, Max: %.4f, Min: %.4f\n", city, avg, stats.Max, stats.Min)
	}

	fmt.Printf("Waktu penyelesaian: %v\n", elapsed)

}

func processChunk(filename string, startOffset int64, endOffset int64, result chan<- map[string]*Data,
	wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error Open File:", err)
		return
	}

	defer file.Close()

	// 1. lompat langsung ke posisi byte awal jatah worker ini
	_, err = file.Seek(startOffset, io.SeekStart)
	if err != nil {
		fmt.Println("Error Seek File:", err)
		return
	}

	reader := bufio.NewReader(file)
	currentPos := startOffset

	// 2. Jika bukan chunk pertama, skip sisa baris parsial sampai ketemu newline `\n`
	if startOffset != 0 {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}
		currentPos += int64(len(line))
	}

	localData := make(map[string]*Data)

	// 3. Baca terus baris per baris selama belum melewati endOffset
	for currentPos < endOffset {
		lineBytes, err := reader.ReadBytes('\n')
		if len(lineBytes) == 0 {
			break
		}

		currentPos += int64(len(lineBytes))

		// ubah byte jadi string dan bersihin newline
		line := strings.TrimSpace(string(lineBytes))
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) < 2 {
			continue
		}

		town := strings.TrimSpace(parts[0])
		temp, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			continue
		}

		stats, exist := localData[town]
		if !exist {
			localData[town] = &Data{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			if temp < stats.Min {
				stats.Min = temp
			}
			if temp > stats.Max {
				stats.Max = temp
			}
			stats.Sum += temp
			stats.Count++
		}

		if err != nil {
			break
		}
	}

	result <- localData

}
