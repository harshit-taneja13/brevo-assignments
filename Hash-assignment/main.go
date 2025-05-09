package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateHash("Harshit TanejaHarshit TanejaHarshit TanejaHarshit TanejaHarshit Taneja")) // 0dCUC9xiw1
	fmt.Println(generateHash("Harshit Taneja"))                                                         // 5Exb2kYSnN
	fmt.Println(generateHash("Harshit Taneja"))                                                         // 5Exb2kYSnN
	fmt.Println(generateHash("Harshit Sangwan"))                                                        // JkuOKdiqIr
	fmt.Println(generateHash("Harshit T"))                                                              // l2sRlGsHzG
	fmt.Println(generateHash("Harshit S"))                                                              // l2qsyRMaLb
	fmt.Println(generateHash("abc"))                                                                    // qGRuIyUwcV
	fmt.Println(generateHash("abd"))                                                                    // qGTT5o0eGA
	fmt.Println(generateHash("b"))                                                                      // 3b2t0IniAn
	fmt.Println(generateHash("a"))                                                                      // 3b1wvntUOC
	fmt.Println(generateHash("a "))                                                                     // jpJXyqkzY4
	fmt.Println(generateHash("aa "))                                                                    // qCOZVDBPr3
	fmt.Println(generateHash("aa"))                                                                     // jpdqDeP8rf
	fmt.Println(generateHash("0000"))                                                                   // iFZcX4Sr4f
	fmt.Println(generateHash("00"))                                                                     // gV9xsN0XBF
	fmt.Println(generateHash("BREVO"))                                                                  // dwvAiPsxd9
	fmt.Println(generateHash("brevo"))                                                                  // fVAIjrgWaF
	fmt.Println(generateHash("hello123"))																// ZjqkDtKapB
	fmt.Println(generateHash("hello12 "))																// Zjuns3GDt6
	fmt.Println(generateHash(" "))																		// 3aiHPKqovf
	// ADDED EMPTY INPUT TEST CASE
	fmt.Println(generateHash(""))																		// Input Cannot be empty : Enter a valid String
																
}
