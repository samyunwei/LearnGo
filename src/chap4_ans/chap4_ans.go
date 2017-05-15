package main

import (
	"fmt"
	"strings"
)

func UniqueInts(input []int) []int {
	res := []int{}
	resmap := map[int]bool{}
	for _, each := range input {
		if _, found := resmap[each]; !found {
			res = append(res, each)
			resmap[each] = true
		}
	}
	return res
}

func Flatten(irregularMatrix [][]int) []int {
	res := []int{}
	for _, eachVector := range irregularMatrix {
		res = append(res, eachVector...)
	}
	return res
}

func Make2D(input []int, sep int) [][]int {
	res := [][]int{}
	beg := 0
	for ; (beg + sep) < len(input); beg += sep {
		res = append(res, input[beg:beg+sep])
	}
	last := input[beg:]
	for ; len(last) < sep; {
		last = append(last, 0)
	}
	return append(res, last)
}

func parseIni(inistrs []string) map[string]map[string]string {
	var lastkey string
	res := map[string]map[string]string{}
	for _, eachline := range inistrs {
		eachline = strings.TrimSpace(eachline)
		if len(eachline) == 0 || (strings.Index(eachline, ";") == 0) {
			continue
		}
		if strings.Index(eachline, "[") == 0 && strings.Index(eachline, "]") == len(eachline)-1 {
			lastkey = eachline[1:len(eachline)-1]
			res[lastkey] = map[string]string{}
			continue
		}
		kv := strings.Split(eachline, "=")
		res[lastkey][kv[0]] = kv[1]
	}
	return res
}

func PrintIni(inidict map[string]map[string]string) {
	for k, v := range inidict {
		fmt.Printf("[%s]\n", k)
		for dk, dv := range v {
			fmt.Printf("%s=%s\n", dk, dv)
		}
	}
}

func main() {
	//testUniqueInts
	//testarray := []int{9,1,9,5,4,4,2,1,5,4,8,8,4,3,6,9,5,7,5}
	//fmt.Println(UniqueInts(testarray))
	//fmt.Println(testarray)

	//testFlatten
	//irregularMatrix := [][]int{{1, 2, 3, 4},
	//	{5,6,7,8},
	//	{9,10,11},
	//	{12,12,14,15},
	//	{16,17,18,19,20}}
	//slice := Flatten(irregularMatrix)
	//fmt.Printf("1x%d: %v\n",len(slice),slice)

	//testMake2D
	//testslice := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	//fmt.Println(testslice)
	//fmt.Println(Make2D(testslice,6))

	//testIniparse
	iniData := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}
	inidict := parseIni(iniData)
	PrintIni(inidict)
}
