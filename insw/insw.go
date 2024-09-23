package insw

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func Exctract(hscode string) {
	endpoint := "https://api.insw.go.id/api-prod-ba/ref/hscode/komoditas?hs_code=" + hscode
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	// Set header
	req.Header.Set("User-Agent", "My-Custom-User-Agent")
	req.Header.Set("accept", "application/json")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("authorization", "Basic aW5zd18yOmJhYzJiYXM2")
	req.Header.Set("Referrer-Policy", "strict-origin-when-cross-origin")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", hscode, err)
		os.Exit(1)
	}

	file, err := os.Create("res/" + hscode + ".json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	file.Write(b)

}

func HsCodes(file *os.File) []string {
	var result []string
	content := bufio.NewScanner(file)
	for content.Scan() {
		if validateHSCode(content.Text()) {
			result = append(result, content.Text())
		}
	}
	return result
}

func validateHSCode(hscode string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, hscode)
	return match
}
