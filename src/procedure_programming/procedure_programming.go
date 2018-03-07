package main

import (
	"fmt"
	"errors"
	"strings"
	"path/filepath"
	"encoding/json"
	"bytes"
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

func main() {
	//typeassert();
	//classicIF();
	//classifier(5,-17.9,"ZIP",nil,true,complex(1,1))
	reverseJson()
}
