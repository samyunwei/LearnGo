package main

import (
	"net/http"
	"fmt"
	"log"
	"strconv"
	"math/cmplx"
	"math"
)

type quadratic_ans struct {
	a  float64 //x^2
	b  float64 //x
	c  float64 //num
	x1 complex128
	x2 complex128
}

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Quadratic Equation Slover</title>
<body><h3>Quadratic Equation Slover</h3>
<p>Sloves equations of the form ax^2 + bx + c</p>`
	form = `<form action="/" method="POST">
<input type="text" name="a" size="30">
<label for="a">x^2 + </label>
<input type="text" name="b" size="30">
<label for="b">x + </label>
<input type="text" name="c" size="30">
<label for="c">-></label>
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

func processRequest(request *http.Request) (float64, float64, float64, bool) {
	a := FindArg(request, "a")
	b := FindArg(request, "b")
	c := FindArg(request, "c")
	if EqualFloat(a, 0, -1) {
		return 0, 0, 0, false
	} else {
		return a, b, c, true
	}
}

func FindArg(request *http.Request, key string) (float64) {
	if slice, found := request.Form[key]; found && len(slice) == 1 {
		if res, err := strconv.ParseFloat(slice[0], 64); err != nil {
			return 0
		} else {
			return res
		}
	}
	return 0
}

func formatQuestion(a float64, b float64, c float64) (ans quadratic_ans) {
	ans.a = a
	ans.b = b
	ans.c = c
	return ans
}

func slove(ans *quadratic_ans) (res *quadratic_ans) {
	a := complex(ans.a, 0)
	b := complex(ans.b, 0)
	c := complex(ans.c, 0)
	Δ := cmplx.Sqrt(cmplx.Pow(b, 2) - (4 * a * c))
	ans.x1 = (-b + Δ) / (2 * a)
	ans.x2 = (-b - Δ) / (2 * a)
	return ans
}

func formatSolutions(ans *quadratic_ans) string {
	res := "<p>"
	res += fmt.Sprintf("% gx^2", ans.a)
	if !EqualFloat(ans.b, 0, -1) {
		res += fmt.Sprintf(" %+gx", ans.b)
	}
	if !EqualFloat(ans.c, 0, -1) {
		res += fmt.Sprintf(" %+g", ans.c)
	}
	if EqualComplex(ans.x1, ans.x2) {
		res += fmt.Sprintf(" ->x =%s", FormatComplex(ans.x1))
	} else {
		res += fmt.Sprintf(" ->x =%s or x =%s", FormatComplex(ans.x1), FormatComplex(ans.x2))
	}
	res += "</p>"
	return res
}

func EqualFloat(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}
	return math.Abs(x-y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
}

func FormatComplex(x complex128) string {
	if EqualFloat(imag(x), 0, -1) {
		return fmt.Sprintf("%g", real(x))
	} else {
		return fmt.Sprintf("%g", x)
	}

}

func EqualComplex(x, y complex128) bool {
	if EqualFloat(real(x), real(y), -1) && EqualFloat(imag(x), imag(y), -1) {
		return true
	}
	return false
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if a, b, c, ok := processRequest(request); ok {
			quadratic := formatQuestion(a, b, c)
			fmt.Fprintf(writer, formatSolutions(slove(&quadratic)))
		} else if EqualFloat(a, 0, -1) && !EqualFloat(b, 0, -1) && !EqualFloat(c, 0, -1) {
			fmt.Fprintf(writer, anError, "a can't not be 0")
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
