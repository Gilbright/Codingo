package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Here is the glory!")

	var arr [4]int
	arr[0] = 35
	arr[1] = 32
	fmt.Printf("%v is of the type: %T\n", arr, arr)

	// slices
	var sli = make([]int, 5)
	sli[3] = 49
	fmt.Println(sli)
	sli = append(sli, 997)
	fmt.Println(sli)

	ma := map[string]int{"three": 3, "four": 4}
	ma["one"] = 1
	fmt.Println(ma)

	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "This is a test about the creation of files")
	file.Close()

}
