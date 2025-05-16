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
	id int;
	done bool;
	desc string
}

func main() {
	displayTitle()
	displayMenu()
	
	var tasks = []Item{};

	for { 
		displayEntrySign()
		tasks = handleCommand(getCommand(getArg()), tasks)
	}
}

func displayEntrySign() {
		fmt.Print(">")
}

func displayDelimiter() {
	fmt.Println("--- --- ---")
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

func handleCommand(command Command, tasks []Item) []Item {
	switch command {
		case Add: {
			fmt.Print("Description: ")
			var desc = getLine()
			tasks = append(tasks, Item{done: false, desc: desc, id: 1})

			fmt.Println(tasks)
			displayDelimiter()
		}
		case List: {
			displayDelimiter()

			for _, item := range tasks {
				fmt.Println("Id:", item.id, "Is Done:", item.done, "Description:", item.desc);
			}

			displayDelimiter()
		}
		case Delete: {
			// var id = getArg()
			displayDelimiter()
		}
		case Complete: {
			// var id = getArg()
			displayDelimiter()
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

	return tasks
}

/*
Description:
Create a simple command-line to-do list application. 
Users can add tasks, view the list of tasks, mark tasks as completed, and delete tasks. 
Store the tasks in a file (like a .txt or .json file) so the list is saved even after the program is closed.
*/