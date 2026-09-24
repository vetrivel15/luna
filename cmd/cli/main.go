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
			register()
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

func register() {
	fmt.Println("Enter userid: ")

	scanner := bufio.NewScanner(os.Stdin)
	userId := strings.TrimSpace(scanner.Text())

	if ok := user.IsExist(userId); !ok {
		fmt.Println("User already exists")
	}

	if ok, err := registeruser(userId); !ok {
		fmt.Println(fmt.Errorf("Error registering user %v", err))
	}

}

func registeruser(id string) (bool, error) {

	fmt.Println("Enter first name:")

	scanner := bufio.NewScanner(os.Stdin)

	firstName := strings.TrimSpace(scanner.Text())

	fmt.Println("Enter last name:")

	scanner = bufio.NewScanner(os.Stdin)

	lastName := strings.TrimSpace(scanner.Text())

	fmt.Println("Enter password")

	scanner = bufio.NewScanner(os.Stdin)

	password := strings.TrimSpace(scanner.Text())

	return user.New(id, firstName, lastName, password)
}

func login() {

}

func deleteUser() {

}
