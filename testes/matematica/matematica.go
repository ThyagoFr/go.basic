package matematica

// Media é responsavel por calcular média '-'
import (
	"fmt"
	"strconv"
)

func Media(numbers ...float64) float64 {
	total := 0.0
	for _,value := range numbers {
		total += value
	}
	mean := total/float64(len(numbers))
	roundedMean,_ := strconv.ParseFloat(fmt.Sprintf("%.2f",mean), 64)
	return roundedMean
}


