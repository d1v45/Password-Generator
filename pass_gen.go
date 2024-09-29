//#####################################################################################################
// Title: PASSWORD GENERATOR                                                                          #
// Author: DIVAS A S                                                                                  #
// Github : https://github.com/d1v45                                                                  #
//#####################################################################################################

package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

func banner() {
	fmt.Println(`
	▒░█▀▀▄ █▀▀█ █▀▀▀ █▀▀▀   █▀▀█ █▀▀▀ █▄  █░▒       
	▒░█▄▄▀░█▄▄█ ▀▀▀█ ▀▀▀█   █ ▄▄ █▀▀░ █▀█▄█░▒       
	▒░█░▒  ▀░░▀ ▀░░▀ ▀░░▀   █▄▄█ ▀░░▀ ▀░░░▀░▒       
					  -d1v45`)
}

func getCharacterSet(includeUpper, includeLower, includeNum, includeSym bool, customSet string) (string, error) {
	var charSet strings.Builder

	if includeUpper {
		charSet.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if includeLower {
		charSet.WriteString("abcdefghijklmnopqrstuvwxyz")
	}
	if includeNum {
		charSet.WriteString("0123456789")
	}
	if includeSym {
		charSet.WriteString("!@#$%^&*()-_=+[]{}|;:,.<>?/`~")
	}
	if customSet != "" {
		charSet.WriteString(customSet)
	}

	if charSet.Len() == 0 {
		return "", errors.New("no character sets selected")
	}

	return charSet.String(), nil
}

func generatePassword(length int, charSet string) (string, error) {
	password := make([]byte, length)

	for i := range password {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}
		password[i] = charSet[index.Int64()]
	}

	return string(password), nil
}

func passGen(length, amount int, charSet string) {
	for i := 0; i < amount; i++ {
		password, err := generatePassword(length, charSet)
		if err != nil {
			fmt.Println("Error generating password:", err)
			return
		}
		fmt.Println(password)
	}
}

func getYesOrNo(prompt string) bool {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scan(&input)

		if strings.ToLower(input) == "y" {
			return true
		} else if strings.ToLower(input) == "n" {
			return false
		} else {
			fmt.Println("Invalid input. Please enter 'y' for yes or 'n' for no.")
		}
	}
}

func main() {
	banner()

	var length, amount int
	var customSet string
	var includeUpper, includeLower, includeNum, includeSym bool

	fmt.Print("Enter the length of the password (minimum 6): ")
	fmt.Scan(&length)
	if length < 6 {
		fmt.Println("Password length should be at least 6.")
		return
	}

	fmt.Print("Enter the number of passwords you need: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Please enter a positive number for the amount.")
		return
	}

	includeUpper = getYesOrNo("Include uppercase letters? (y/n): ")
	includeLower = getYesOrNo("Include lowercase letters? (y/n): ")
	includeNum = getYesOrNo("Include numbers? (y/n): ")
	includeSym = getYesOrNo("Include symbols? (y/n): ")

	if includeUpper || includeLower || includeNum || includeSym {
		if getYesOrNo("Would you like to enter a custom character set? (y/n): ") {
			fmt.Print("Enter a custom character set: ")
			fmt.Scan(&customSet)
		}
	}

	charSet, err := getCharacterSet(includeUpper, includeLower, includeNum, includeSym, customSet)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	passGen(length, amount, charSet)
}
