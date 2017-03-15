package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
	"math"
	"bytes"
	"io"
	"strconv"
	"regexp"
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

func SimpleSimplifyWhitespace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

func SimpleSimplifyWhitespace2(s string) string {
	var buffer bytes.Buffer
	skip := true
	for _, char := range s {
		if unicode.IsSpace(char) {
			if !skip {
				buffer.WriteRune(' ')
				skip = true
			}
		} else {
			buffer.WriteRune(char)
			skip = false
		}
	}
	s = buffer.String()
	if skip && len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

func StringBagstrings() {
	names := "ddd,aasd,dadxq,aada,,ada,azza,aa"
	fmt.Print("|")
	for _, name := range strings.SplitAfter(names, ",") {
		fmt.Printf("%s|", name)
	}
	fmt.Println()

	for _, record := range []string{"aaaa ssss*121212*2222", "daaaa cadada\t111\t11sss", "sdad cxccc gdgfg|1775|1814"} {
		fmt.Println(strings.FieldsFunc(record, func(char rune) bool {
			switch char {
			case '\t', '*', '|':
				return true
			}
			return false
		}))
	}

	names = "adada\tdada\tadaa\t\t\tcaxaxaxa\t\tdadaaadad\tldllff"
	names = strings.Replace(names, "\t", " ", -1)
	fmt.Printf("|%s|\n", names)
	fmt.Printf("|%s|\n", SimpleSimplifyWhitespace2(names))

	asciiOnly := func(char rune) rune {
		if char > 127 {
			return '?'
		}
		return char
	}
	fmt.Println(strings.Map(asciiOnly, "哈哈哈zzzzddd中文"))

	reader := strings.NewReader("daa哈迭代aa")
	for {
		char, size, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Printf("%U %c %d: % X\n", char, char, size, []byte(string(char)))
	}
}

func Stringstrconv() {
	for _, truth := range []string{"1", "t", "TRUE", "false", "F", "0", "5"} {
		if b, err := strconv.ParseBool(truth); err != nil {
			fmt.Printf("\n{%v}", err)
		} else {
			fmt.Print(b, " ")
		}
	}
	fmt.Println()

	x, err := strconv.ParseFloat("-99.7", 64)
	fmt.Printf("%8T %6v %v\n", x, x, err)
	y, err := strconv.ParseInt("71309", 10, 0)
	fmt.Printf("%8T %6v %v\n", y, y, err)
	z, err := strconv.Atoi("71309")
	fmt.Printf("%8T %6v %v\n", z, z, err)

	s := strconv.FormatBool(z > 100)
	fmt.Println(s)
	i, err := strconv.ParseInt("0xDEED", 0, 32)
	fmt.Println(i, err)
	j, err := strconv.ParseInt("0707", 0, 32)
	fmt.Println(j, err)
	k, err := strconv.ParseInt("10111010001", 2, 32)
	fmt.Println(k, err)

	l := 16769023
	fmt.Println(strconv.Itoa(l))
	fmt.Println(strconv.FormatInt(int64(l), 10))
	fmt.Println(strconv.FormatInt(int64(l), 2))
	fmt.Println(strconv.FormatInt(int64(l), 16))

	s = "这是中文 aaaa ddaad ddaaq vvva aa.\n"
	quoted := strconv.Quote(s)
	fmt.Println(quoted)
	fmt.Println(strconv.Unquote(quoted))

}

func IsHexDigit(char rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, char)
}

func StringFmtunicode() {
	fmt.Println(IsHexDigit('8'), IsHexDigit('x'), IsHexDigit('X'), IsHexDigit('b'), IsHexDigit('B'))
}

func StringRegxReplace() {
	names := []string{"aaaa dddd zzz aaa dddd cadaadd"}
	println("Regx1")
	nameRx := regexp.MustCompile(`(\pL+\.?(?:\s+\pL+\.?)*)\s+(\pL+)`)
	for i := 0; i < len(names); i++ {
		names[i] = nameRx.ReplaceAllString(names[i], "${2},${1}")
		println(names[i])
	}

	names = []string{"aaaa dddd zzz aaa dddd cadaadd"}
	println("Regx2")
	nameRx2 := regexp.MustCompile(`(?P<forename>\pL+\.?(?:\s+\pL+\.?)*)\s+(?P<surename>\pL+)`)
	for i := 0; i < len(names); i++ {
		names[i] = nameRx2.ReplaceAllString(names[i], "${surename} ${forename}")
		println(names[i])
	}
}

func StringRegxMatchWord() {
	text := "dada aaa aaa cca dada ccc aaa xxx"
	wordRx := regexp.MustCompile(`\w+`)
	if mathches := wordRx.FindAllString(text, -1); mathches != nil {
		previous := ""
		for _, match := range mathches {
			if match == previous {
				fmt.Println("Duplicate word:", match)
			}
			previous = match
		}
	}
}

func StringRegxMap() {
	lines := " a : vv"
	valueForKey := make(map[string]string)
	keyValueRx := regexp.MustCompile(`\s*([[:alpha:]]\w*)\s* :\s*(.+)`)
	if matches := keyValueRx.FindAllStringSubmatch(lines, -1); matches != nil {
		for _, match := range matches {
			valueForKey[match[1]] = strings.TrimRight(match[2], `\t`)
		}
		fmt.Println(valueForKey)
	}

}

func StringRegxXML() {
	attribs := `a="name"`
	atttName := "a"
	attrValueRx := regexp.MustCompile(regexp.QuoteMeta(atttName) + `=(?:"([^"]+)"|'([^']+)')`)
	if indexes := attrValueRx.FindAllStringSubmatchIndex(attribs, -1); indexes != nil {
		for _, positions := range indexes {
			start, end := positions[2], positions[3]
			if start == -1 {
				start, end = positions[4], positions[5]
			}
			fmt.Printf("`%s`\n", attribs[start:end])
		}
	} else {
		println("match error")
	}
}

func StringRegxSimplifyWhitespace() {
	text := "aaa       ddd        "
	SimpleSimplifyWhitespaceRx := regexp.MustCompile(`[\s\p{Zl}\p{Zp}]+`)
	text = strings.TrimSpace(SimpleSimplifyWhitespaceRx.ReplaceAllLiteralString(text, " "))
	println(text)
}

func StringRegxReplaceByfunc() {
	text := "aaaa bbbbb"
	replacerx := regexp.MustCompile(`a`)
	text = replacerx.ReplaceAllStringFunc(text, func(s string) string {
		return string("s")
	})
	println(text)
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
	//StringFmtDebug()
	//StringBagstrings()
	//Stringstrconv()
	//StringFmtunicode()
	//StringRegxReplace()
	//StringRegxMatchWord()
	//StringRegxMap()
	//StringRegxXML()
	//StringRegxSimplifyWhitespace()
	StringRegxReplaceByfunc()
}
