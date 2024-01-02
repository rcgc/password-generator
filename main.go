package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	lowerCharSet   	= "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   	= "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet 	= "!@#$%&*"
	numberSet	   	= "0123456789"
	allCharSet	   	= lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	// Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := r.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	// Set numeric
	for i := 0; i < minNum; i++ {
		random := r.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	// Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := r.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := r.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int){
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func main() {
	password := generatePassword(20, 1, 1, 1)
	fmt.Println(password)
}