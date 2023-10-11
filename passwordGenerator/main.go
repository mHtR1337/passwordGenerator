package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()_+{}[]|//;:'<>,.?/~"
)

func GeneratePassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	allChars := uppercaseLetters + lowercaseLetters + digits + specialChars

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	return string(password)
}

func main() {
	var passwordLength int

	fmt.Print("Enter an password length: ")
	_, err := fmt.Scanf("%d", &passwordLength)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	password := GeneratePassword(passwordLength)
	fmt.Println("Generated Password: ", password)
}
