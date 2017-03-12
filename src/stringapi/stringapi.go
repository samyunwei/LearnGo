package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

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

func main() {
	//stringoperators()
	//charAndString()
	//RuneandIndex()
	//StringIndex()
	StringIndex2()
}
