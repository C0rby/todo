package todofile

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type TodoFile struct {
	Name string
	Path string
}

func CreateName() string {
	return time.Now().UTC().Format("02-01-2006")
}

func New(name, path string) *TodoFile {
	return &TodoFile{Name: name, Path: path}
}

func (f *TodoFile) Exists() bool {
	_, err := os.Stat(f.Path)
	return !os.IsNotExist(err)
}

func (f *TodoFile) Add(todo string) {
	file, err := os.OpenFile(f.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0644))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(todo + "\n")
}

func (f *TodoFile) Read() string {
	b, err := ioutil.ReadFile(f.Path)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (f *TodoFile) ReadLines(lines int) string {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var buffer bytes.Buffer

	scanner := bufio.NewScanner(file)
	for i := 1; i <= lines && scanner.Scan(); i++ {
		buffer.WriteString(scanner.Text())
		buffer.WriteByte('\n')
	}
	return buffer.String()
}
