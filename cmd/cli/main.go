package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	user "github.com/vetrivel15/luna/internal/model"
)

func main() {

	fmt.Println("Enter your choice.\n1: Register\n2: Login\n3: Deactivate user\nQ: Quit")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {

		input := strings.TrimSpace(scanner.Text())
		switch strings.ToLower(input) {

		case "1":
			register(scanner)
		case "2":
			login()
		case "3":
			deleteUser()
		case "q":
		default:
			fmt.Println("Invalid choice")
		}

	}
}

func register(scanner *bufio.Scanner) {
	fmt.Println("Enter userid: ")

	scanner.Scan()
	userId := strings.TrimSpace(scanner.Text())

	if ok := user.IsExist(userId); ok {
		fmt.Println("User already exists")
	}

	if ok, err := registeruser(scanner, userId); !ok {
		fmt.Println(fmt.Errorf("Error registering user %v", err))
	}

	fmt.Println("User registeration complete")

}

func registeruser(scanner *bufio.Scanner, id string) (bool, error) {

	fmt.Println("Enter first name:")

	scanner.Scan()

	firstName := strings.TrimSpace(scanner.Text())

	fmt.Println("Enter last name:")

	scanner.Scan()

	lastName := strings.TrimSpace(scanner.Text())

	fmt.Println("Enter password")

	scanner.Scan()

	password := strings.TrimSpace(scanner.Text())

	return user.New(id, firstName, lastName, password)
}

func login() {

}

func deleteUser() {

}
