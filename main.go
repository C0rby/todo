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
	todo ls [<folder>]
	todo mkdir <folder>
	todo rm <folder>
	todo head [<folder>]
	todo cat [<folder>]
	todo -h | --help
	todo <todo>
	todo <folder> <todo>
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
		todo.Add(folder, arguments["<todo>"].(string))
	} else if arguments["cat"] == true {
		todo.ReadCurrent(folder)
	} else if arguments["head"] == true {
		todo.ReadLinesCurrent(folder, 5)
	}

	//fmt.Println(arguments)
}
