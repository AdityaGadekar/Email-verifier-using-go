package main

import (
    "fmt"
    "regexp"
)

func main() {
	var email string
    fmt.Println("Enter email to verify")
	fmt.Scanln(&email)
    fmt.Println(email,"is\t:",isEmailValid(email)) // true
}

func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return emailRegex.MatchString(e)
}

