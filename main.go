package main

import (
	"fmt"
	"strings"
	"os"
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

type Item struct {
	done bool;
	desc string
}

func main() {
	displayTitle()
	displayMenu()
	
	for { 
		handleCommand(getCommand(getArg())) 
	}
}

func displayTitle() {
	fmt.Println("- -- -- -- -")
	fmt.Println("* Todo App *")
	fmt.Println("- -- -- -- -")
}

func displayMenu() {
	fmt.Println("Commands:")
	fmt.Println(" Add - add new task")
	fmt.Println(" List - list all tasks")
	fmt.Println(" Delete <Task Id> - delete task")
	fmt.Println(" Complete <Task Id> - mark task as completed")
	fmt.Println(" Quit - exit program")
	fmt.Println(" Help - show this message")
}


func getCommand(word string) Command {
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

func getArg() string {
	var word string
	fmt.Scan(&word)
	return word
}

func getLine() string {
	var word string
	fmt.Scanln(&word)
	return word
}

func handleCommand(command Command) {
	switch command {
		case Add: {
			fmt.Println("Description: ")
			var desc = getLine()
		}
		case List: {
			
		}
		case Delete: {
			var id = getArg()
		}
		case Complete: {
			var id = getArg()
		}
		case Help: {
			displayMenu()
		}
		case None: {
			fmt.Println("Unrecognized command.")
			displayMenu()
		}
		case Quit: {
			os.Exit(0)
		}
	}
}

/*
Description:
Create a simple command-line to-do list application. 
Users can add tasks, view the list of tasks, mark tasks as completed, and delete tasks. 
Store the tasks in a file (like a .txt or .json file) so the list is saved even after the program is closed.
*/