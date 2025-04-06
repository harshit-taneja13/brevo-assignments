package main

import "fmt"

func customHash(input string) uint64 {
	
	var hash uint64 = 14695981039346656037
	const prime uint64 = 1099511628211

	for _, ch := range input {
		hash ^= uint64(ch)
		hash *= prime
	}
	return hash
}

func base62Conversion(num uint64) string {
	const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if num == 0 {
		return "0"
	}
	encoded := ""
	base := uint64(62)

	for num > 0 {
		remainder := num % base
		encoded = string(base62Chars[remainder]) + encoded
		num /= base
	}
	return encoded
}

const mod uint64 = 62 * 62 * 62 * 62 * 62 * 62 * 62 * 62 * 62 * 62

func generateHash(input string) string {
	hashVal := customHash(input)

	reduced := hashVal % mod

	hashString := base62Conversion(reduced)

	for len(hashString) < 10 {
		hashString = "0" + hashString
	}

	return hashString
}


func main() {
	fmt.Println(generateHash("Harshit TanejaHarshit TanejaHarshit TanejaHarshit TanejaHarshit Taneja"))
	fmt.Println(generateHash("Harshit Taneja"))
	fmt.Println(generateHash("Harshit Taneja"))
	fmt.Println(generateHash("Harshit Sangwan"))
	fmt.Println(generateHash("Harshit T"))
	fmt.Println(generateHash("Harshit S"))
	fmt.Println(generateHash("abc"))
	fmt.Println(generateHash("abd"))
	fmt.Println(generateHash("b"))
	fmt.Println(generateHash("a"))
}
