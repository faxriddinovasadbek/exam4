package etc

import (
	// "fmt"
	"math/rand"
	"time"
)

// Table for code generator
var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// Generate code is function that create n-digit random code
func GenerateCode(max int) string {
	rand.Seed(time.Now().UnixNano()) // Rand funktsiyasining nima uchun istifoda qilinayotgan vaqtini sozlash
	b := make([]byte, max)
	for i := range b {
		b[i] = table[rand.Intn(len(table))]
	}
	return string(b)
}

func GenerateStrongPasswordWithDigits(max int) string {
	length := max
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const digits = "0123456789"
	const totalLetters = len(letters)
	const totalDigits = len(digits)

	var password string

	for i := 0; i < length/2; i++ {
		password += string(letters[rand.Intn(totalLetters)])
	}

	for i := 0; i < length/2; i++ {
		password += string(digits[rand.Intn(totalDigits)])
	}

	runePassword := []rune(password)
	rand.Shuffle(len(runePassword), func(i, j int) {
		runePassword[i], runePassword[j] = runePassword[j], runePassword[i]
	})
	password = string(runePassword)

	return password
}
