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

func (t *Todo) Add(folder, todo string) {
	fileName := todofile.CreateName()
	filePath := filepath.Join(t.conf.BaseDir, folder, fileName)
	f := todofile.New(fileName, filePath)
	f.Add(todo)
}

func (t *Todo) Read(folder, name string) {
	filePath := filepath.Join(t.conf.BaseDir, folder, name)
	f := todofile.New(name, filePath)
	fmt.Print(f.Read())
}

func (t *Todo) ReadCurrent(folder string) {
	t.Read(folder, todofile.CreateName())
}

func (t *Todo) ReadLinesCurrent(folder string, lines int) {
	fileName := todofile.CreateName()
	filePath := filepath.Join(t.conf.BaseDir, folder, fileName)
	f := todofile.New(fileName, filePath)
	fmt.Print(f.ReadLines(lines))
}

func (t *Todo) Undo(folder, todo string) {
	fileName := todofile.CreateName()
	todoPath := filepath.Join(t.conf.BaseDir, folder, fileName)
	fTodo := todofile.New(fileName, todoPath)
	fTodo.DeleteLines(todo)
}

func (t *Todo) Complete(folder, todo string) {
	t.Undo(folder, todo)

	doneName := todofile.CreateName() + "-done"
	donePath := filepath.Join(t.conf.BaseDir, folder, doneName)
	fDone := todofile.New(doneName, donePath)
	fDone.Add(todo)
}
