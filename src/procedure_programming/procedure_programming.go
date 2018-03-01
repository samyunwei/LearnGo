package main

import (
	"fmt"
	"errors"
	"strings"
	"path/filepath"
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

func main() {
	//typeassert();
	classicIF();
}
