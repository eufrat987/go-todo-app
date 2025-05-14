package main

import (
	"fmt"
	"strings"
)

type Command int
const (
	Add Command = iota
	List
	Delete
	Complete
	Quit
	Help
	None
)

func displayMenu() {
	fmt.Println("Commands:")
	fmt.Println(" Add - add new task")
	fmt.Println(" List - list all tasks")
	fmt.Println(" Delete <Task Id> - delete task")
	fmt.Println(" Complete <Task Id> - mark task as completed")
	fmt.Println(" Quit - exit program")
	fmt.Println(" Help - show this message")
}

func getCommand() Command {
	var word string
	fmt.Scan(&word)
	
	switch strings.ToLower(word) {
		case "add": return Add
		case "list": return List
		case "delete": return Delete
		case "complete": return Complete
		case "quit": return Quit
		case "help": return Help
		default: return None
	}
}

func getArgIfNeeded(command Command) string {
	if command == Delete || command == Complete {
		var word string
		fmt.Scan(&word)
		return word
	}
	
	return ""
}

func main() {
	fmt.Println("- -- -- -- -")
	fmt.Println("* Todo App *")
	fmt.Println("- -- -- -- -")
	
	displayMenu()
	
	for {
		
		var command = getCommand()
		// var arg = getArgIfNeeded(command)
		
		if command == Quit {
			break
		}
	}
}

/*
Description:
Create a simple command-line to-do list application. 
Users can add tasks, view the list of tasks, mark tasks as completed, and delete tasks. 
Store the tasks in a file (like a .txt or .json file) so the list is saved even after the program is closed.
*/