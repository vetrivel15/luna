package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	id        string
	firstName string
	lastName  string
	password  string
}

func main() {

	fmt.Println("Enter your choice.\n1: Register\n2: Login\n3: Deactivate user\nQ: Quit")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {

		input := strings.TrimSpace(scanner.Text())
		switch strings.ToLower(input) {

		case "1":
			registerNewUser()
		case "2":
			loginUser()
		case "3":
			deregisterUser()
		case "q":
		default:
			fmt.Println("Invalid choice")
		}

	}
}

func registerNewUser() bool {
	fmt.Println("Enter userid: ")

	scanner := bufio.NewScanner(os.Stdin)

	userId := strings.TrimSpace(scanner.Text())

	if isUserExist(userId) {
		fmt.Println("User already exists")
		return false
	}

	registerUser()

	return true
}

func isUserExist(id string) bool {
	return false
}
