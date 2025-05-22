package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./app <nama> <usia>")
		return
	}

	name := os.Args[1]
	age, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Usia harus berupa angka")
		return
	}

	fmt.Printf("Halo, %s! Usia kamu adalah %d tahun.\n", name, age)
}
