package main

import (
	"fmt"
	"errors"
	"strings"
	"path/filepath"
	"encoding/json"
	"bytes"
	"math/rand"
	"os"
	"log"
	"math"
	"net/http"
)

func typeassert() {
	var i interface{} = 99
	var s interface{} = []string{"left", "right"}
	j := i.(int)
	fmt.Printf("%T->%d\n", j, j)
	if i, ok := i.(int); ok {
		fmt.Printf("%T->%d\n", i, j)
	}
	if s, ok := s.([]string); ok {
		fmt.Printf("%T->%s\n", s, s)
	}
}

func classicIF() {
	if a := -1; a < 0 {
		fmt.Printf("(%d)\n", -a)
	} else {
		fmt.Println(a)
	}
}

func GzipFileList(file string) ([]string, error) {
	return nil, errors.New("Gzipfiles")
}

func TarFileList(file string) ([]string, error) {
	return nil, errors.New("TarFileList")
}

func ZipFileList(file string) ([]string, error) {
	return nil, errors.New("Zipfiles")
}

func Suffix(file string) string {
	file = strings.ToLower(filepath.Base(file))
	if i := strings.LastIndex(file, "."); i > -1 {
		if file[i:] == ".bz2" || file[i:] == ".gz" || file[i:] == ".xz" {
			if j := strings.LastIndex(file[:i], ".");
				j > -1 && strings.HasPrefix(file[j:], ".tar") {
				return file[j:]
			}
		}
		return file[:i]
	}
	return file
}

func ArchiveFileList(file string) ([]string, error) {
	if suffix := Suffix(file); suffix == ".gz" {
		return GzipFileList(file)
	} else if suffix == ".tar" || suffix == ".tar.gz" || suffix == ".gz" {
		return TarFileList(file)
	} else if suffix == ".zip" {
		return ZipFileList(file)
	}
	return nil, errors.New("unrecognized archive")
}

func BundedInt(minimum, value, maximum int) int {
	switch {
	case value < minimum:
		return minimum
	case value > maximum:
		return maximum
	}
	return value
}

func BundedInt2(minimum, value, maximum int) int {
	switch {
	case value < minimum:
		return minimum
	case value > maximum:
		return maximum
	default:
		return value
	}
	panic("unreachable")
}

func ArchiveFileList2(file string) ([]string, error) {
	switch suffix := Suffix(file); suffix {
	case ".gz":
		return GzipFileList(file)
	case ".tar":
		fallthrough
	case ".tar.gz":
		fallthrough
	case ".tgz":
		return TarFileList(file)
	case ".zip":
		return ZipFileList(file)
	}
	return nil, errors.New("unrecognized archive")
}

func ArchiveFileList3(file string) ([]string, error) {
	switch suffix := Suffix(file); suffix {
	case ".gz", ".tar", ".tar.gz":
		return GzipFileList(file)
	case ".tgz":
		return TarFileList(file)
	case ".zip":
		return ZipFileList(file)
	}
	return nil, errors.New("unrecognized archive")
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int8, int16, int32, int64:
			fmt.Printf("param #%d is an int\n", i)
		case uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is an unsigned int\n", i)
		case nil:
			fmt.Printf("param #%d is a nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d is unknow\n", i)
		}
	}
}

func JsonObjectAsString(jsonObject map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	comma := ""
	for key, value := range jsonObject {
		buffer.WriteString(comma)
		switch value := value.(type) {
		case nil:
			fmt.Fprintf(&buffer, "%q: null", key)
		case bool:
			fmt.Fprintf(&buffer, "%q: %t", key, value)
		case float64:
			fmt.Fprintf(&buffer, "%q: %f", key, value)
		case string:
			fmt.Fprintf(&buffer, "%q: %q", key, value)
		case []interface{}:
			fmt.Fprintf(&buffer, "%q: [", key)
			innerComma := ""
			for _, s := range value {
				if s, ok := s.(string); ok {
					fmt.Fprintf(&buffer, "%s%q", innerComma, s)
					innerComma = ", "
				}
			}
			buffer.WriteString("]")
		}
		comma = ", "
	}
	buffer.WriteString("}")
	return buffer.String()
}

func reverseJson() {
	MA := []byte(`{"name": "Massachusetts", "area":27336, "water":25.7,"senators":["John Kerry","Scott Brown"]}`)
	var object interface{}
	if err := json.Unmarshal(MA, &object); err != nil {
		fmt.Println(err)
		fmt.Println("error!!")
	} else {
		jsonObject := object.(map[string]interface{})
		fmt.Println(JsonObjectAsString(jsonObject))
	}
}

type State struct {
	Name     string
	Senators []string
	Water    float64
	Area     int
}

func (state State) String() string {
	var senators []string
	for _, senator := range state.Senators {
		senators = append(senators, fmt.Sprintf("%q", senator))
	}
	return fmt.Sprintf(
		`{"name": %q, "area":%d, "water":%f,"senators":[%s]}`, state.Name, state.Area, state.Water, strings.Join(senators, ", "))
}

func reverseJson2() {
	var state State
	MA := []byte(`{"name": "Massachusetts", "area":27336, "water":25.7,"senators":["John Kerry","Scott Brown"]}`)
	if err := json.Unmarshal(MA, &state); err != nil {
		fmt.Println(err)
	}
	fmt.Println(state)
}

func testfor(table [][]int, x int) {
	found := false
	for row := range table {
		for column := range table[row] {
			if table[row][column] == x {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}

func testfor2(table [][]int, x int) {
FOUND:
	for row := range table {
		for column := range table[row] {
			if table[row][column] == x {
				break FOUND
			}
		}
	}
}

func createCounter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)
	return next
}

func testchannel() {
	counterA := createCounter(2)
	counterB := createCounter(102)
	for i := 0; i < 5; i++ {
		a := <-counterA
		fmt.Printf("(A-> %d,B ->%d)", a, <-counterB)
	}
	fmt.Println()
}
func testselect() {
	channels := make([]chan bool, 6)
	for i := range channels {
		channels[i] = make(chan bool)
	}
	go func() {
		for {
			channelid := rand.Intn(6)
			channels[channelid] <- true
			//fmt.Printf("channelid:%d\n",channelid)
		}
	}()

	for i := 0; i < 36; i++ {
		var x int
		select {
		case <-channels[0]:
			x = 1
		case <-channels[1]:
			x = 2
		case <-channels[2]:
			x = 3
		case <-channels[3]:
			x = 4
		case <-channels[4]:
			x = 5
		case <-channels[5]:
			x = 6
		}
		fmt.Printf("%d", x)
	}
	fmt.Println()
}

type Data struct {
}

func expensiveComputation(data Data, answer chan int, done chan bool) {
	finished := false
	for !finished {
		result := 2
		answer <- result
	}
	done <- true
}

func testexpensivecomputation() {
	const allDone = 2
	doneCount := 0
	answera := make(chan int)
	answerb := make(chan int)
	defer func() {
		close(answera)
		close(answerb)
	}()
	done := make(chan bool)
	defer func() {
		close(done)
	}()
	var data1 Data
	var data2 Data
	go expensiveComputation(data1, answera, done)
	go expensiveComputation(data2, answerb, done)
	for doneCount != allDone {
		var which, result int
		select {
		case result = <-answera:
			which = 'a'
		case result = <-answerb:
			which = 'b'
		case <-done:
			doneCount++
		}
		if which != 0 {
			fmt.Printf("%c -> %d", which, result)
		}
	}
	fmt.Println()
}
func testdefer(filename string) {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file", err)
		return
	}
	defer file.Close()
}

func ConvertInt64ToInt(x int64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		return int(x)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", x))
}

func IntFromInt64(x int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	i = ConvertInt64ToInt(x)
	return i, nil
}
func homepage(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
		}
	}()
}

func testdefer2() {
	http.HandleFunc("/", homepage)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func logPanics(function func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		function(writer, request)
	}
}

func main() {
	//typeassert();
	//classicIF();
	//classifier(5,-17.9,"ZIP",nil,true,complex(1,1))
	//reverseJson()
	//reverseJson2()
	//testchannel()
	testselect()
}
