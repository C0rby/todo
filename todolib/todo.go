package todolib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/c0rby/todo/todolib/config"
	"github.com/c0rby/todo/todolib/todofile"
)

type Todo struct {
	conf config.Config
}

func New(conf config.Config) *Todo {
	return &Todo{conf}
}

func (t *Todo) BaseDirExists() bool {
	_, err := os.Stat(t.conf.BaseDir)
	return !os.IsNotExist(err)
}

func (t *Todo) CreateBaseDir() {
	t.MakeDir("")
}

func (t *Todo) ListFiles(path string) {
	files, err := ioutil.ReadDir(filepath.Join(t.conf.BaseDir, path))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Print("+ ")
		} else {
			fmt.Print("- ")
		}
		fmt.Println(f.Name())
	}
}

func (t *Todo) MakeDir(name string) {
	if err := os.Mkdir(filepath.Join(t.conf.BaseDir, name), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func (t *Todo) RemoveDir(name string) {
	if err := os.Remove(filepath.Join(t.conf.BaseDir, name)); err != nil {
		log.Fatal(err)
	}
}

func (t *Todo) Add(folder string, todo string) {
	fileName := todofile.CreateName()
	filePath := filepath.Join(t.conf.BaseDir, folder, fileName)
	fmt.Println(filePath)
	f := todofile.New(fileName, filePath)
	f.Add(todo)
}

func (t *Todo) ReadCurrent(folder string) {
	fileName := todofile.CreateName()
	filePath := filepath.Join(t.conf.BaseDir, folder, fileName)
	f := todofile.New(fileName, filePath)
	fmt.Print(f.Read())
}
