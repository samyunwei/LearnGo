package main

import (
	"strings"
	"strconv"
	"log"
	"path/filepath"
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
)

type Song struct {
	Title    string
	Filename string
	Seconds  int
}

func parseExtinfLine(line string) (title string, seconds int) {
	if i := strings.IndexAny(line, "-0123456789"); i > -1 {
		const separator = ","
		line = line[i:]
		if j := strings.Index(line, separator); j > -1 {
			title = line[j+len(separator):]
			var err error
			if seconds, err = strconv.Atoi(line[:j]); err != nil {
				log.Printf("failed to read the duration for `%s`: %v\n", title, err)
				seconds = -1
			}
		}
	}
	return title, seconds
}

func mapPlarformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}

func readM3uPlaylist(data string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			song.Title, song.Seconds = parseExtinfLine(line)
		} else {
			song.Filename = strings.Map(mapPlarformDirSeparator, line)
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

func readPlsPlaylist(data string) (songs []Song) {
	var song Song
	typematch := regexp.MustCompile(`(?P<type>\w+)=(?P<attrval>.+)`)
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "[playlist]" {
			continue
		}
		if res := typematch.FindStringSubmatch(line); res != nil {
			if strings.HasPrefix(res[1], "File") {
				song.Filename = res[2]
			} else if strings.HasPrefix(res[1], "Title") {
				song.Title = res[2]
			} else {
				if seconds, err := strconv.Atoi(res[2]); err != nil {
					log.Printf("failed to read the duration for `%s`: %v\n", song.Title, err)
					seconds = -1
				} else {
					song.Seconds = seconds
				}
			}
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

func writePlsPlaylist(songs []Song) {
	fmt.Println("[playlist]")
	for i, song := range songs {
		i++
		fmt.Printf("File%d=%s\n", i, song.Filename)
		fmt.Printf("Title%d=%s\n", i, song.Title)
		fmt.Printf("Length%d=%d\n", i, song.Seconds)
	}
	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}

func writem3ulist(songs []Song) {
	fmt.Print("#EXTM3U\n")
	for i, song := range songs {
		i++
		fmt.Printf("#EXTINF:%d,%s\n", song.Seconds, song.Title)
		fmt.Printf("%s\n", song.Filename)
	}
	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}

func main() {
	if len(os.Args) == 1 || !(strings.HasSuffix(os.Args[1], `.m3u`) || strings.HasSuffix(os.Args[1], `.pls`)) {
		fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(os.Args[0]))
		os.Exit(0)
	} else {

	}
	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else if strings.HasSuffix(os.Args[1], `.m3u or .pls`) {

		songs := readM3uPlaylist(string(rawBytes))
		writePlsPlaylist(songs)
	} else {
		songs := readPlsPlaylist(string(rawBytes))
		writem3ulist(songs)
	}
}
