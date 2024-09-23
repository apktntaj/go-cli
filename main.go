package main

import (
	"fmt"
	"os"
	"time"

	"github.com/apktntaj/go-cli/insw"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func main() {
	file, err := os.Open("hs_code_indonesia.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	defer file.Close()
	hsCodes := insw.HsCodes(file)

	start := time.Now()
	for _, hsCode := range hsCodes {
		insw.Exctract(hsCode)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
