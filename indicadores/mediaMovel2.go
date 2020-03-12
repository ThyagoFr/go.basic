package main

import "fmt"

func main(){
	values := []float64{10,20,30,50,60,70,80,90,100}
	averages := []float64{}
	n := 3
	sum := 0.0
	for i:= 0; i < n; i++ {
		sum += values[i]
	}
	average := sum/float64(n)
	averages = append(averages,average)
	for j := n; j < len(values); j++ {
		average = (average*float64(n) - values[j-n] + values[j])/float64(n)
		averages = append(averages,average)
	}
	fmt.Println(averages)
}
