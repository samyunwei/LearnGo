package main

import (
	"fmt"
	"image/color"
	"strings"
	"sort"
)

type composer struct {
	name      string
	birthYear int
}

type rectangle struct {
	x0, y0, x1, y1 int
	fill           color.RGBA
}

type Product struct {
	name  string
	price float64
}

type Point struct {
	x, y, z int
}

func (point Point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", point.x, point.y, point.z)
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

func (product Product) String() string {
	return fmt.Sprintf("%s (%.2f)", product.name, product.price)
}

func InsertStringSliceCopy(slice, insetion []string, index int) []string {
	result := make([]string, len(slice)+len(insetion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insetion)
	copy(result[at:], slice[index:])
	return result
}

func InsertStringSLice(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}

func RemoveStringSliceCopy(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}

func RemoveStringSlice(slice []string, start, end int) []string {
	return append(slice[:start], slice[end:]...)
}

func SortFoldedStrings(slice []string) {
	sort.Sort(FoldedStrings(slice))
}

type FoldedStrings []string

func (slice FoldedStrings) Len() int {
	return len(slice)
}

func (slice FoldedStrings) Less(i, j int) bool {
	return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}

func (slice FoldedStrings) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
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

func arrayapi() {
	var buffer [20]byte
	var grid1 [3][3]int
	grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
	grid2 := [3][3]int{{4, 3}, {8, 6, 2}}
	citites := [...]string{"Shanghai", "Mumbau", "Istanbul", "Beijing"}
	citites[len(citites)-1] = "Karachi"
	fmt.Println("Type Len Contents")
	fmt.Printf("%-8T %2d %v\n", buffer, len(buffer), buffer)
	fmt.Printf("%-8T %2d %v\n", citites, len(citites), citites)
	fmt.Printf("%-8T %2d %v\n", grid1, len(grid1), grid1)
	fmt.Printf("%-8T %2d %v\n", grid2, len(grid2), grid2)

}

func sliceapi1() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := s[:5]
	u := s[3:len(s)-1]
	fmt.Print(s, t, u)
	u[1] = "X"
	fmt.Println(s, t, u)
}

func sliceapi2() {
	s := new([7]string)[:]
	s[0], s[1], s[2], s[3], s[4], s[5], s[6] = "A", "B", "C", "D", "E", "F", "G"
	buffer := make([]byte, 20, 60)
	grid1 := make([][]int, 3)
	for i := range grid1 {
		grid1[i] = make([]int, 3)
	}
	grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
	grid2 := [][]int{{4, 3}, {8, 6, 2}, {0, 0, 0}}
	citites := []string{"Shanghai", "Mumbau", "Istanbul", "Beijing"}
	citites[len(citites)-1] = "Karachi"
	fmt.Println("Type Len Contents")
	fmt.Printf("%-8T %2d %v\n", buffer, len(buffer), buffer)
	fmt.Printf("%-8T %2d %v\n", citites, len(citites), citites)
	fmt.Printf("%-8T %2d %v\n", grid1, len(grid1), grid1)
	fmt.Printf("%-8T %2d %v\n", grid2, len(grid2), grid2)
}

func sliceapi3() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := s[2:6]
	fmt.Println(t, s, "=", s[:4], "+", s[4:])
	s[3] = "x"
	t[len(t)-1] = "y"
	fmt.Println(t, s, "=", s[:4], "+", s[4:])
}

func sliceapi4() {
	amounts := []float64{237.81, 261.87, 273.93, 279.99, 281.07, 303.17, 231.47, 227.33, 209.23, 197.09}
	sum := 0.0
	for _, amount := range amounts {
		sum += amount
	}
	fmt.Printf("∑ %.1f -> %.1f\n", amounts, sum)
}

func sliceapi5() {
	amounts := []float64{237.81, 261.87, 273.93, 279.99, 281.07, 303.17, 231.47, 227.33, 209.23, 197.09}
	sum := 0.0
	for i := range amounts {
		amounts[i] *= 1.05
		sum += amounts[i]
	}
	fmt.Printf("∑ %.1f -> %.1f\n", amounts, sum)
}

func sliceapi6() {
	products := []*Product{{"Spanner", 3.99}, {"Wrench", 2.49}, {"Screwdirver", 1.99}}
	fmt.Println(products)
	for _, product := range products {
		product.price += 0.50
	}
	fmt.Println(products)

}

func sliceapi7() {
	s := []string{"A", "B", "C", "D", "E", "F"}
	t := []string{"K", "L", "M", "N"}
	u := []string{"m", "n", "o", "p", "q", "r"}
	s = append(s, "h", "i", "j")
	s = append(s, t...)
	s = append(s, u[2:5]...)
	b := []byte{'U', 'V'}
	letters := "XWY"
	b = append(b, letters...)
	fmt.Printf("%v\n%s\n", s, b)
}

func sliceapi8() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	x := InsertStringSliceCopy(s, []string{"a", "b", "c"}, 0)
	y := InsertStringSliceCopy(s, []string{"x", "y"}, 3)
	z := InsertStringSliceCopy(s, []string{"z"}, len(s))
	fmt.Printf("%v\n%v\n%v\n%v\n", s, x, y, z)
}

func sliceapi9() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	s = s[2:]
	fmt.Println(s)
}

func sliceapi10() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	s = s[:4]
	fmt.Println(s)
}

func sliceapi11() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	s = append(s[:1], s[5:]...)
	fmt.Println(s)
}

func sliceapi12() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	x := RemoveStringSliceCopy(s, 0, 2)
	y := RemoveStringSliceCopy(s, 1, 5)
	z := RemoveStringSliceCopy(s, 4, len(s))
	fmt.Printf("%v\n%v\n%v\n%v\n", s, x, y, z)

}

func sliceapi13() {
	files := []string{"Test.conf", "uitl.go", "Makefile", "misc.go", "main.go"}
	fmt.Printf("Unsorted:          %q\n", files)
	sort.Strings(files)
	fmt.Printf("Underlying bytes:          %q\n", files)
	SortFoldedStrings(files)
	fmt.Printf("Case insentitive:          %q\n", files)
}

func sliceapi14() {
	files := []string{"Test.conf", "uitl.go", "Makefile", "misc.go", "main.go"}
	target := "Makefile"
	for i, file := range files {
		if file == target {
			fmt.Printf("found \"%s\" at files[%d]\n", file, i)
			break
		}
	}
}

func sliceapi15() {
	files := []string{"Test.conf", "uitl.go", "Makefile", "misc.go", "main.go"}
	target := "Makefile"
	sort.Strings(files)
	fmt.Printf("%q\n", files)
	i := sort.Search(len(files), func(i int) bool { return files[i] >= target })
	if i < len(files) && files[i] == target {
		fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
	}
}

func sliceapi16() {
	files := []string{"Test.conf", "uitl.go", "Makefile", "misc.go", "main.go"}
	target := "makefile"
	SortFoldedStrings(files)
	fmt.Printf("%q\n", files)
	caseInsensitiveCampare := func(i int) bool { return strings.ToLower(files[i]) >= target }
	i := sort.Search(len(files), caseInsensitiveCampare)
	if i <= len(files) && strings.EqualFold(files[i], target) {
		fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
	}
}

func mapapi() {
	massforPlanet := make(map[string]float64)
	massforPlanet["Mercurt"] = 0.06
	massforPlanet["Venus"] = 0.82
	massforPlanet["Earth"] = 1.00
	massforPlanet["Mars"] = 0.11
	fmt.Println(massforPlanet)
}

func mapapi2() {
	triangle := make(map[*Point]string, 3)
	triangle[&Point{89, 47, 27}] = "a"
	triangle[&Point{86, 65, 86}] = "b"
	triangle[&Point{7, 44, 45}] = "r"
	fmt.Println(triangle)
}

func mapapi3() {
	triangle := make(map[Point]string, 3)
	triangle[Point{89, 47, 27}] = "x"
	triangle[Point{86, 65, 86}] = "y"
	fmt.Println(triangle)
}

func mapapi4() {
	populationForCity := map[string]int{"Istanbul": 12610000, "Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
	for city, popluation := range populationForCity {
		fmt.Printf("%-10s %8d\n", city, popluation)
	}
}

func mapapi5() {
	populationForCity := map[string]int{"Istanbul": 12610000, "Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
	popluation := populationForCity["Mumbai"]
	fmt.Println("Mumbai's poplation is ", popluation)
	popluation = populationForCity["Emerald City"]
	fmt.Println("Emerald City's poplation is ", popluation)

}

func mapapi6() {
	populationForCity := map[string]int{"Istanbul": 12610000, "Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
	city := "Istanbul"
	if population, found := populationForCity[city]; found {
		fmt.Printf("%s's population is %d \n", city, population)
	} else {
		fmt.Printf("%s's popluation data is unavailable\n", city)
	}
	city = "Emerald City"
	_, prsent := populationForCity[city]
	fmt.Printf("%q is in the map == %t\n", city, prsent)
}

func mapapi7() {
	populationForCity := map[string]int{"Istanbul": 12610000, "Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
	fmt.Println(len(populationForCity), populationForCity)
	delete(populationForCity, "Shanghai")
	fmt.Println(len(populationForCity), populationForCity)
	populationForCity["Karachi"] = 11620000
	fmt.Println(len(populationForCity), populationForCity)
	populationForCity["Beijing"] = 11290000
	fmt.Println(len(populationForCity), populationForCity)
}

func mapapi8() {
	populationForCity := map[string]int{"Istanbul": 12610000, "Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
	oldKey, newKey := "Beijing", "Tokyo"
	value := populationForCity[oldKey]
	delete(populationForCity, oldKey)
	populationForCity[newKey] = value + 1
	fmt.Println(len(populationForCity), populationForCity)
}

func mapapi9() {
	populationForCity := map[string]int{"Istanbul": 12610000, "Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
	cities := make([]string, 0, len(populationForCity))
	for city := range populationForCity {
		cities = append(cities, city)
	}
	sort.Strings(cities)
	for _, city := range cities {
		fmt.Printf("%-10s %8d\n", city, populationForCity[city])
	}

}

func mapapi10() {
	populationForCity := map[string]int{"Istanbul": 12610000, "Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
	citiesForPopulation := make(map[int]string, len(populationForCity))
	for city, population := range populationForCity {
		citiesForPopulation[population] = city
	}
	fmt.Println(citiesForPopulation)
}

func main() {
	//ptrapi1()
	//ptrapi2()
	//ptrapi4()
	//ptrapi5()
	//arrayapi()
	//sliceapi1()
	//sliceapi2()
	//sliceapi3()
	//sliceapi4()
	//sliceapi5()
	//sliceapi6()
	//sliceapi7()
	//sliceapi8()
	//sliceapi9()
	//sliceapi10()
	//sliceapi11()
	//sliceapi12()
	//sliceapi13()
	//sliceapi14()
	//sliceapi15()
	//sliceapi16()
	//mapapi()
	//mapapi2()
	//mapapi3()
	//mapapi4()
	//mapapi5()
	//mapapi6()
	//mapapi7()
	//mapapi8()
	//mapapi9()
	mapapi10()
}
