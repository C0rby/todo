package main

import (
	"log"

	"github.com/c0rby/todo/todolib"
	cfg "github.com/c0rby/todo/todolib/config"
	"github.com/docopt/docopt-go"
)

const configPath = ".config/todo/config.json"

var (
	config *cfg.Config
	todo   *todolib.Todo
)

func init() {
	config = cfg.Load(configPath)
	todo = todolib.New(*config)

	if !todo.BaseDirExists() {
		todo.CreateBaseDir()
	}
}

func main() {
	usage := `
Usage:
	todo (<folder> ls| ls)
	todo mkdir <folder>
	todo rm <folder>
	todo (<folder> head| head)
	todo (<folder> cat [<name>]| cat [<name>])
	todo (<folder> [(-d|--done)] <todo>| [(-d|--done)] <todo>)
	todo (<folder> [(-u|--undo)] <todo>| [(-u|--undo)] <todo>)
	todo -h | --help
	todo --version
	`
	arguments, err := docopt.Parse(usage, nil, true, "todo 0.1", false)
	if err != nil {
		log.Fatal("uh oh" + err.Error())
	}

	var folder string
	if arguments["<folder>"] != nil {
		folder = arguments["<folder>"].(string)
	}

	if arguments["ls"] == true {
		todo.ListFiles(folder)
	} else if arguments["mkdir"] == true {
		todo.MakeDir(arguments["<folder>"].(string))
	} else if arguments["rm"] == true {
		todo.RemoveDir(arguments["<folder>"].(string))
	} else if arguments["<todo>"] != nil {
		t := arguments["<todo>"].(string)
		if arguments["-d"] == true || arguments["--done"] == true {
			todo.Complete(folder, t)
			return
		} else if arguments["-u"] == true || arguments["--undo"] == true {
			todo.Undo(folder, t)
			return
		}
		todo.Add(folder, t)
	} else if arguments["cat"] == true {
		if arguments["<name>"] == nil {
			todo.ReadCurrent(folder)
		} else {
			todo.Read(folder, arguments["<name>"].(string))
		}
	} else if arguments["head"] == true {
		todo.ReadLinesCurrent(folder, 5)
	}
	//fmt.Println(arguments)
}
