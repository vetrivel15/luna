package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	user "github.com/vetrivel15/luna/internal/model"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nEnter your choice")
		fmt.Println("1: Register")
		fmt.Println("2: Login")
		fmt.Println("3: Deactivate user")
		fmt.Println("Q: Quit")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		switch strings.ToLower(input) {

		case "1":
			register(scanner)
		case "2":
			login(scanner)
		case "3":
			deleteUser(scanner)
		case "q":
			fmt.Println("Exiting application...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func register(scanner *bufio.Scanner) {
	fmt.Println("Enter userid: ")

	scanner.Scan()
	userId := strings.TrimSpace(scanner.Text())

	fmt.Println("Enter first name:")

	scanner.Scan()

	firstName := strings.TrimSpace(scanner.Text())

	fmt.Println("Enter last name:")

	scanner.Scan()

	lastName := strings.TrimSpace(scanner.Text())

	fmt.Println("Enter password")

	scanner.Scan()

	password := strings.TrimSpace(scanner.Text())

	if ok, err := user.New(userId, firstName, lastName, password); !ok {
		fmt.Printf("Failed user registeration %v", err)
		return
	}

	fmt.Println("User registeration complete")
}

func login(scanner *bufio.Scanner) {
	fmt.Println("\nLogin")
	fmt.Println("Enter userid: ")

	scanner.Scan()
	userId := strings.TrimSpace(scanner.Text())

	fmt.Println("Enter password")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	if ok, err := user.Login(userId, password); !ok {
		fmt.Printf("User login failed: %v", err)
		return
	}

	fmt.Println("User login complete")
}

func deleteUser(scanner *bufio.Scanner) {
	fmt.Printf("\nDelete user")

	fmt.Println("Enter userid: ")
	scanner.Scan()
	userid := strings.TrimSpace(scanner.Text())

	fmt.Println("Enter password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	if ok, err := user.Delete(userid, password); !ok {
		fmt.Printf("Delete user failed: %v", err)
		return
	}

	fmt.Println("Delete user complete")
}
