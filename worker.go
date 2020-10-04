package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readFile(fname string) string {
	fileHandle, err := os.Open(fname)

	if err != nil {
		panic(err)
	}

	defer fileHandle.Close()

	line := ""
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	for {
		cursor--
		fileHandle.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		fileHandle.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	return line
}

func addLineToFile(fname string, line string) (string, *bool) {
	f, err := os.OpenFile(fname,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
		return "", new(bool)
	}

	defer f.Close()

	if _, err := f.WriteString("\n" + line); err != nil {
		log.Println(err)
		return "", new(bool)
	}

	return line, nil
}

func getLatestTag(t *Tag) (string, *bool) {

	tag := readFile(t.file)

	if tag == "" {
		panic("No tag found")
	}

	return tag, nil
}
