package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		// io.Copy(os.Stdout, resp.Body) Cara alternatif tanpa malalui proses pembacaan byte
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		fmt.Printf(resp.Status)
	}
}
