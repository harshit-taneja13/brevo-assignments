package main

const (
	prime uint64		 = 1099511628211    // FNV-1a 64-bit prime multiplier (chosen experimentally by checking collision rate)
	FnvOffset          	 = 14695981039346656037
    Base62Chars          = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    Base                 = 62
    MaxBase62HashValue   = 62 * 62 * 62 * 62 * 62 * 62 * 62 * 62 * 62 * 62  // this is a constant representing 62^10, which is the maximum value that can be represented by a 10-character Base62 string.
)
