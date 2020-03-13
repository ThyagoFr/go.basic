package main

import "fmt"

func main() {
	values := []float64{22,30,40.7,57,69}
	averages := []float64{}
	n := 2
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += values[i]
	}
	average := sum / float64(n)
	averages = append(averages, values[:n]...)
	averages = append(averages, average)
	for j := n; j < len(values); j++ {
		average = (average*float64(n) - values[j-n] + values[j]) / float64(n)
		averages = append(averages, average)
	}
	fmt.Println(averages)
}
