package main

import "fmt"

func main() {
	
	fmt.Println(generateHash("Harshit TanejaHarshit TanejaHarshit TanejaHarshit TanejaHarshit Taneja"))  // 0dCUC9xiw1
	fmt.Println(generateHash("Harshit Taneja"))   // 5Exb2kYSnN
	fmt.Println(generateHash("Harshit Taneja"))   // 5Exb2kYSnN
	fmt.Println(generateHash("Harshit Sangwan"))  // JkuOKdiqIr
	fmt.Println(generateHash("Harshit T"))        // l2sRlGsHzG
	fmt.Println(generateHash("Harshit S"))        // l2qsyRMaLb	
	fmt.Println(generateHash("abc"))              // qGRuIyUwcV
	fmt.Println(generateHash("abd"))              // qGTT5o0eGA
	fmt.Println(generateHash("b"))                // 3b2t0IniAn
	fmt.Println(generateHash("a"))                // 3b1wvntUOC

}
