package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile := ""
	flag.StringVar(&inputFile, "input", "", "The input file containing the station measurement data")
	flag.Parse()

	if inputFile == "" {
		panic("No measurements file provided")
	}

	stations := map[string]*stationData{}

	file, err := os.OpenFile(inputFile, os.O_RDONLY, 0644)
	if err != nil {
		panic("Failed to open file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		lineParts := strings.Split(fileScanner.Text(), ";")

		if _, found := stations[lineParts[0]]; !found {
			stations[lineParts[0]] = NewStationData()
		}

		station := stations[lineParts[0]]
		temp, _ := strconv.ParseFloat(lineParts[1], 64) // Generated data assumed to be valid

		if temp > station.high {
			station.high = temp
		}
		if temp < station.low {
			station.low = temp
		}
		station.total += temp
		station.readCount++
	}

	stationNames := make([]string, 0, len(stations))
	for stationName := range stations {
		stationNames = append(stationNames, stationName)
	}

	sort.Strings(stationNames)

	output := "{"

	for _, stationName := range stationNames {
		station := stations[stationName]
		output += fmt.Sprintf("%s=%.1f/%.1f/%.1f, ", stationName, station.low, station.total/float64(station.readCount), station.high)
	}

	output = strings.TrimRight(output, ", ") + "}"
	fmt.Println(output)
}

type stationData struct {
	high      float64
	low       float64
	total     float64
	readCount int
}

func NewStationData() *stationData {
	return &stationData{
		high:      math.Inf(-1),
		low:       math.Inf(1),
		total:     0,
		readCount: 0,
	}
}
