# todo
A command line todo tool

## Introduction

todo is a simple command line utility for managing (obviously) todos. The todos are stored in files that are either located in the home directory under .todo/ or at a path specified in the configuration file.

## Usage
```bash
Usage:
        todo (<folder> ls| ls)
        todo mkdir <folder>
        todo rm <folder>
        todo (<folder> head | head)
        todo (<folder> cat [<name>] | cat [<name>])
        todo (<folder> <todo> | <todo>)
        todo (<folder> (-d | --done) <todo>| (-d | --done) <todo>)
        todo (<folder> (-u | --undo) <todo>| (-u | --undo) <todo>)
        todo -h | --help
        todo --version
```
