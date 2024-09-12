package main

import (
	"fmt"
	"os"

	"github.com/apktntaj/go-cli/insw"
)

func main() {
	file, err := os.Open("bup copy.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	defer file.Close()

	result := insw.HsCodes(file)

	fmt.Println(result)

}
