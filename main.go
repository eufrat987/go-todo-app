package main

import "fmt"

func main() {
	fmt.Println("- -- -- -- -")
	fmt.Println("* Todo App *")
	fmt.Println("- -- -- -- -")
	
	fmt.Println("Commands:")
	fmt.Println(" Add - add new task")
	fmt.Println(" List - list all tasks")
	fmt.Println(" Delete <Task Id> - delete task")
	fmt.Println(" Complete <Task Id> - mark task as completed")
	fmt.Println(" Quit - exit program")
	fmt.Println(" Help - show this message")

	for {
		break
	}
}

/*
Description:
Create a simple command-line to-do list application. 
Users can add tasks, view the list of tasks, mark tasks as completed, and delete tasks. 
Store the tasks in a file (like a .txt or .json file) so the list is saved even after the program is closed.
*/