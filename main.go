package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readDataFromFile(filePath string) ([]float64, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var data []float64

	for _, line := range lines {
		if line != "" {
			value, err := strconv.ParseFloat(line, 64)
			if err != nil {
				return nil, err
			}
			data = append(data, value)
		}
	}

	return data, nil
}

func calculateAverage(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func calculateMedian(data []float64) float64 {
	sort.Float64s(data)
	length := len(data)

	if length%2 == 0 {
		mid1 := data[length/2-1]
		mid2 := data[length/2]
		return (mid1 + mid2) / 2
	}

	return data[length/2]
}

func calculateVariance(data []float64, average float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += math.Pow(value-average, 2)
	}
	return sum / float64(len(data))
}

func calculateStandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run your-program.go data.txt")
	}

	filePath := os.Args[1]

	data, err := readDataFromFile(filePath)
	if err != nil {
		log.Fatal("Error reading data from file:", err)
	}

	average := calculateAverage(data)
	median := calculateMedian(data)
	variance := calculateVariance(data, average)
	standardDeviation := calculateStandardDeviation(variance)

	fmt.Printf("Average: %d\n", int(math.Round(average)))
	fmt.Printf("Median: %d\n", int(math.Round(median)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(standardDeviation)))
}
