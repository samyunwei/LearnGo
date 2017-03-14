package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
	"math"
)

type polar struct {
	radius, o float64
}

func stringoperators() {
	book := "The Sprit Level" + "By Richard Wilkinson"
	book += " and Kate Pickett"
	fmt.Println(book)
	fmt.Println("Josay" < "Jose", "Josey" == "Jose")
}

func charAndString() {
	s := ""
	for _, char := range []rune{'æ', 0xE6, 0346, 230, '\xE6', '\u00e6'} {
		fmt.Printf("[0x%X '%c']", char, char)
		s += string(char)
	}
	fmt.Printf(string('\n') + s)
}

func RuneandIndex() {
	phrase := "this is string to parse"
	fmt.Printf("string: \"%s\"\n", phrase)
	fmt.Println("index rune char bytes")
	for index, char := range phrase {
		fmt.Printf("%-2d    %U  '%c' %X\n", index, char, char, []byte(string(char)))
	}

}

func StringIndex() {
	line := "r哈哈ca你 ccc 猜几个串"
	i := strings.Index(line, " ")
	fistWord := line[:i]
	j := strings.LastIndex(line, " ")
	lastword := line[j+1:]
	fmt.Println(fistWord, lastword)
}

func StringIndex2() {
	line := "r 啊啊啊啊啊啊啊 aaaaa \u2028r谢谢"
	i := strings.IndexFunc(line, unicode.IsSpace)
	firstword := line[:i]
	j := strings.LastIndexFunc(line, unicode.IsSpace)
	_, size := utf8.DecodeRuneInString(line[j:])
	lastword := line[j+size:]
	fmt.Println(firstword, lastword)
}

func StringFmtPrint() {

	p := polar{8.32, .49}
	fmt.Print(-18.5, 17, "Elephant", -8+.7i, 0x3c7, '\u03C7', "a", "b", p)
	fmt.Println()
	fmt.Println(-18.5, 17, "Elephant", -8+.7i, 0x3c7, '\u03C7', "a", "b", p)

}

func IntForBool(b bool) int {
	if b {
		return 1
	}
	return 0
}
func StringFmtBool() {
	fmt.Printf("%t %t\n", true, false)
	fmt.Printf("%d %d\n", IntForBool(true), IntForBool(false))
}
func Pad(number, width int, pad rune) string {
	s := fmt.Sprint(number)
	gap := width - utf8.RuneCountInString(s)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + s
	}
	return s
}
func StringFmtInt() {
	fmt.Printf("|%b|%9b|%-9b|%09b|% 09b|\n", 37, 37, 37, 37, 37)
	fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41, 41, 41, 41, -41)
	i := 3931
	fmt.Printf("|%x|%X|%8x|%08x|%#04X|0x%04X|\n", i, i, i, i, i, i)
	i = 569
	fmt.Printf("|$%d|$%06d|$%+06d|$%s|\n", i, i, i, Pad(i, 6, '*'))
}

func StringFmtString() {
	fmt.Printf("%d %#04x %U '%c' \n", 0x3A6, 934, '\u03A6', '\U000003A6')
}

func Humanize(amount float64, width, decimals int, pad, separator rune) string {
	dollars, cents := math.Modf(amount)
	whole := fmt.Sprintf("%+.0f", dollars)[1:]
	fraction := ""
	if decimals > 0 {
		fraction = fmt.Sprintf("%+.*f", decimals, cents)[2:]
	}
	sep := string(separator)
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + sep + whole[i:]
	}
	if amount < 0.0 {
		whole = "-" + whole
	}
	number := whole + fraction
	gap := width - utf8.RuneCountInString(number)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + number
	}
	return number
}

func StringFmtFloat() {
	for _, x := range [] float64{-.258, 7194.84, -60897162.0218, 1.500089e-8} {
		fmt.Printf("|%20.5e|%20.5f|%s|\n", x, x, Humanize(x, 20, 5, '*', ','))
	}

	for _, x := range []complex128{2 + 3i, 172.6 - 58.3019i, -.827e2 + 9.04831e-3i} {
		fmt.Printf("|%15s|%9.3f|%.2f|%.1e|\n", fmt.Sprintf("%6.2f%+.3fi", real(x), imag(x)), x, x, x)
	}

}

func StringFmtSlice() {
	slogn := "beg 你好iiii哈哈"
	fmt.Printf("%s\n%q\n%+q\n%#q\n", slogn, slogn, slogn, slogn)

	chars := []rune(slogn)
	fmt.Printf("%x\n%#x\n%#X\n", chars, chars, chars)

	bytes := []byte(slogn)
	fmt.Printf("%s\n%x\n%X\n% X\n%v\n", bytes, bytes, bytes, bytes, bytes)

	s := "Dare to be naive"
	fmt.Printf("|%22s|%-22s|%10s|\n", s, s, s);
	i := strings.Index(s, "n")
	fmt.Printf("|%.10s|%.*s|%-22.10s|%s|\n", s, i, s, s, s)
}

func StringFmtDebug() {
	p := polar{-83.40, 71.60}
	fmt.Printf("|%T|%v|%#v|\n", p, p, p)
	fmt.Printf("|%T|%v|%t|\n", false, false, false)
	fmt.Printf("|%T|%v|%d|\n", 7607, 7607, 7607)
	fmt.Printf("|%T|%v|%f|\n", math.E, math.E, math.E)
	fmt.Printf("|%T|%v|%f|\n", 5+7i, 5+7i, 5+7i)

	s := "Relativity"
	fmt.Printf("|%T|\"%v\"\"%s\"|%q|\n", s, s, s, s)

	s = "Alias是Synonym"
	chars := []rune(s)
	bytes := []byte(s)
	fmt.Printf("%T: %v\n%T: %v\n", chars, chars, bytes, bytes)

	i := 5
	f := -48.3124
	s = "this is string"
	fmt.Printf("|%p -> %d|%p -> %f|%#p -> %s\n", &i, i, &f, f, &s, s)

	fmt.Println([]float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%v\n", []float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%#v\n", []float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%.5f\n", []float64{math.E, math.Pi, math.Phi})

	fmt.Printf("%q\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%v\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%#v\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%17s\n", []string{"Software patents", "kill", "innovation"})

	fmt.Printf("%v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
	fmt.Printf("%#v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
	fmt.Printf("%v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
	fmt.Printf("%#v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
	fmt.Printf("%04b\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})

}

func main() {
	//stringoperators()
	//charAndString()
	//RuneandIndex()
	//StringIndex()
	//StringIndex2()
	//StringFmtPrint()
	//StringFmtBool()
	//StringFmtInt()
	//StringFmtString()
	//StringFmtFloat()
	//StringFmtSlice()
	StringFmtDebug()
}
