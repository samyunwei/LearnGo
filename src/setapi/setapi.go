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

func arrayapi()  {
	var buffer [20]byte
	var grid1 [3][3]int
	grid1[1][0],grid1[1][1],grid1[1][2] = 8,6,2
	grid2 := [3][3]int{{4,3},{8,6,2}}
	citites := [...]string{"Shanghai","Mumbau","Istanbul","Beijing"}
	citites[len(citites)-1] = "Karachi"
	fmt.Println("Type Len Contents")
	fmt.Printf("%-8T %2d %v\n",buffer,len(buffer),buffer)
	fmt.Printf("%-8T %2d %v\n",citites,len(citites),citites)
	fmt.Printf("%-8T %2d %v\n",grid1,len(grid1),grid1)
	fmt.Printf("%-8T %2d %v\n",grid2,len(grid2),grid2)

}

func sliceapi1()  {
	s := []string{"A","B","C","D","E","F","G"}
	t := s[:5]
	u := s[3:len(s)-1]
	fmt.Print(s,t,u)
	u[1] = "X"
	fmt.Println(s,t,u)
}

func sliceapi2(){
	s := new([7]string)[:]
	s[0],s[1],s[2],s[3],s[4],s[5],s[6] = "A","B","C","D","E","F","G"
	buffer := make([]byte,20,60)
	grid1 := make([][]int,3)
	for i := range grid1{
		grid1[i] = make([]int,3)
	}
	grid1[1][0],grid1[1][1],grid1[1][2] = 8,6,2
	grid2 := [][]int{{4,3},{8,6,2},{0,0,0}}
	citites := []string{"Shanghai","Mumbau","Istanbul","Beijing"}
	citites[len(citites)-1] = "Karachi"
	fmt.Println("Type Len Contents")
	fmt.Printf("%-8T %2d %v\n",buffer,len(buffer),buffer)
	fmt.Printf("%-8T %2d %v\n",citites,len(citites),citites)
	fmt.Printf("%-8T %2d %v\n",grid1,len(grid1),grid1)
	fmt.Printf("%-8T %2d %v\n",grid2,len(grid2),grid2)
}

func sliceapi3(){
	s := []string{"A","B","C","D","E","F","G"}
	t := s[2:6]
	fmt.Println(t,s,"=",s[:4],"+",s[4:])
	s[3] = "x"
	t[len(t)-1] = "y"
	fmt.Println(t,s,"=",s[:4],"+",s[4:])
}

func main() {
	//ptrapi1()
	//ptrapi2()
	//ptrapi4()
	//ptrapi5()
	//arrayapi()
	//sliceapi1()
	//sliceapi2()
	sliceapi3()
}


