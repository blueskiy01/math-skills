package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func readSequence() []float64 {
	var sequence []float64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}
		sequence = append(sequence, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}

	return sequence
}

func calculateVariance(data []float64) float64 {
	average := calculateAverage(data)
	sumOfSquares := 0.0
	for _, value := range data {
		sumOfSquares += math.Pow(value-average, 2)
	}
	return sumOfSquares / float64(len(data))
}

func calculateAverage(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func printRangeForNextNumber(lastNumber float64, variance float64) {
	// We can choose a range based on the standard deviation.
	// For example, we can consider one standard deviation from the mean.
	// The range can be defined as mean ± standard deviation.
	standardDeviation := math.Sqrt(variance)
	lowerLimit := int(math.Round(lastNumber - standardDeviation))
	upperLimit := int(math.Round(lastNumber + standardDeviation))

	fmt.Printf("%d %d\n", lowerLimit, upperLimit)
}

func main() {
	fmt.Println("Enter the sequence of numbers (press Enter after each number, press Enter twice to finish):")
	sequence := readSequence()

	if len(sequence) == 0 {
		fmt.Println("No input provided.")
		return
	}

	variance := calculateVariance(sequence)
	lastNumber := sequence[len(sequence)-1]

	fmt.Println("Next number range prediction:")
	printRangeForNextNumber(lastNumber, variance)
}
