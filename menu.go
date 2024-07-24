package main

import "fmt"

/*
This function will print out the menu of options on the command line for the user
@return: None
*/
func printMenu() {
	fmt.Println("----- Employe Management System -----")
	fmt.Println("1. Create Employee Record")
	fmt.Println("2. View Employee Record(s)")
	fmt.Println("3. Update Employee Record")
	fmt.Println("4. Delete Employee Record")
	fmt.Println("5. Exit")
	fmt.Println("-------------------------------------")
}
