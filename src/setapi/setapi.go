package main

import (
	"fmt"
	"image/color"
)

type composer struct {
	name      string
	birthYear int
}

type rectangle struct {
	x0, y0, x1, y1 int
	fill           color.RGBA
}

func swapAndProduct1(x, y, product *int) {
	if *x > *y {
		*x, *y = *y, *x
	}
	*product = *x * *y
}

func swapAndProduct2(x, y int) (int, int, int) {
	if x > y {
		x, y = y, x
	}
	return x, y, x * y
}

func inflate(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}

func resizeRect(rect *rectangle, width, height int) {
	(*rect).x1 += width
	rect.y1 += height
}

func ptrapi1() {
	z := 37
	pi := &z
	ppi := &pi
	fmt.Println(z, *pi, **ppi)
	**ppi++
	fmt.Println(z, *pi, **ppi)
}

func ptrapi2() {
	i := 9
	j := 5
	product := 0
	swapAndProduct1(&i, &j, &product)
	fmt.Println(i, j, product)
}

func ptrapi3() {
	i := 9
	j := 5
	i, j, product := swapAndProduct2(i, j)
	fmt.Println(i, j, product)
}

func ptrapi4() {
	antonio := composer{"Antonio Teixeira", 107}
	agnes := new(composer)
	agnes.name, agnes.birthYear = "Agnes Zimmermann", 1845
	julia := &composer{}
	julia.name, julia.birthYear = "Julia Ward Howe", 1819
	augusta := &composer{"Augusta Holmes", 1847}
	fmt.Println(antonio)
	fmt.Println(agnes, augusta, julia)
}

func ptrapi5() {
	grades := []int{87, 55, 43, 71, 60, 43, 21, 19, 63}
	inflate(grades, 3)
	fmt.Println(grades)
}

func ptrapi6() {
	rect := rectangle{4, 8, 20, 10, color.RGBA{0xFF, 0, 0, 0xFF}}
	fmt.Println(rect)
	resizeRect(&rect, 5, 5)
	fmt.Println(rect)
}

func main() {
	//ptrapi1()
	//ptrapi2()
	//ptrapi4()
	ptrapi5()
}
