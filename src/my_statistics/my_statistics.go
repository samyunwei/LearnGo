package main

import (
	"sort"
	"net/http"
	"strings"
	"strconv"
	"fmt"
	"log"
	"math"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
	mode    []float64
	std_dev float64
}

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

func sum(numbers []float64) (total float64) {
	total = 0
	for _, val := range numbers {
		total += val
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result

}

func mod(numbers []float64) ([]float64, bool) {
	numbermap := make(map[float64]int)
	for _, num := range numbers {
		if val, found := numbermap[num]; found {
			numbermap[num] = val + 1
		} else {
			numbermap[num] = 1
		}
	}
	diff := false
	max := 0
	for _, hit := range numbermap {
		if hit > max {
			if max != 0 {
				diff = true
			}
			max = hit
		}
	}
	var res []float64
	for key, val := range numbermap {
		if val == max {
			res = append(res, key)
		}
	}
	return res, diff
}

func std_dev(numbers []float64) (float64, bool) {
	n := len(numbers)
	if n < 1 {
		return -1, false
	}
	avr := sum(numbers) / float64(n)
	var total float64
	for _, num := range numbers {
		total += math.Pow(num-avr, 2)
	}
	return math.Sqrt(total / float64(n-1)), true
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	if res, has := mod(numbers); has {
		stats.mode = res
	}
	if res, has := std_dev(numbers); has {
		stats.std_dev = res
	}
	return stats
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}

	if len(numbers) == 0 {
		return numbers, "", false
	}
	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
<tr><td>Mod</td><td>%v</td></tr>
<tr><td>Std.Dev.</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.mode, stats.std_dev)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprintf(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprintf(writer, pageBottom)
}

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}
