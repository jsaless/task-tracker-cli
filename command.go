package main

import (
	"flag"
	"fmt"
	"os"
)

type CmdFlags struct {
	Add        string
	Del        int
	Update     int
	InProgress int
	Done       int
	List       string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new task specifying a title")
	flag.IntVar(&cf.Del, "delete", -1, "Delete a task by Id")
	flag.IntVar(&cf.Update, "update", -1, "Update a task by Id and specify a new description")
	flag.IntVar(&cf.InProgress, "mark-in-progress", -1, "Update the status to 'in-progress' of a task by Id")
	flag.IntVar(&cf.Done, "mark-done", -1, "Update the status to 'done' of a task by Id")
	flag.StringVar(&cf.List, "list", "", "List tasks by their status")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(tasks *Tasks) {
	switch {
	case cf.Add != "":
		tasks.add(cf.Add)
	case cf.Del != -1:
		tasks.delete(cf.Del)
	case cf.Update != -1:
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli -update <id> <description>")
			return
		}

		description := os.Args[3]

		tasks.update(cf.Update, description)

	case cf.InProgress != -1:
		tasks.status(cf.InProgress, "in-progress")
	case cf.Done != -1:
		tasks.status(cf.Done, "done")
	case cf.List != "":
		tasks.list(cf.List)
	}
}
