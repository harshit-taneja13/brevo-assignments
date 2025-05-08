package main

// customHash computes a 64-bit hash for the specified input string using the FNV-1a algorithm.
// Input - a string that you want to hash.
// Returns - A uint64 value representing the computed hash.
// It only computes and returns the hash.
func customHash(input string) uint64 {
	var hash uint64 = FnvOffset  // FNV-1a 64-bit offset basis (chosen experimentally by checking collision rate)

	for _, ch := range input {
		hash ^= uint64(ch) // To incorporate the input character into the hash
		hash *= prime      // Multiplying by the prime number for dispersion.
	}
	return hash
}

// base62Conversion converts the given unsigned 64-bit number to a Base62 string representation.
// Input - the uint64 number to convert.
// Returns - A string representing 'num' in Base62, using the characters 0-9, A-Z, and a-z.
func base62Conversion(num uint64) string {

	// base62Chars represents the character set for Base62 encoding.
	// It includes digits 0-9, uppercase letters A-Z, and lowercase letters a-z.
	const Base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const Base = uint64(62)

	if num == 0 {
		return "0"
	}

	// encoded stores base 62 converted string
	encoded := ""

	for num > 0 {
		remainder := num % Base
		// appending the corresponding character
		encoded = string(Base62Chars[remainder]) + encoded
		num /= Base
	}
	return encoded
}

// generateHash returns a unique 10-character hash string for the given input.
// The process involves computing a FNV-1a hash of the input, reducing it modulo 62^10,
// converting the result to a Base62 string, and padding it with leading zeros if necessary.
// Input - the input string to hash.
// Returns - A 10-character string representing the hash.
func generateHash(input string) string {
	if len(input) == 0 {
		return "Input Cannot be empty : Enter a valid String"
	}

	// computing FNV-1a hash of the input
	hashVal := customHash(input)
	// reducing it modulo 62^10
	reduced := hashVal % MaxBase62HashValue
	// converting the result to a Base62 string
	hashString := base62Conversion(reduced)

	// padding it with leading zeros if length of string less than 10
	for len(hashString) < 10 {
		hashString = "0" + hashString
	}

	return hashString
}
