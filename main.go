package main

import (
	"bufio"
	"fmt"
	"github.com/hni/sitemap/lib"
	"os"
)

func main() {
	var urls []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
    bytes := sitemap.Sitemap(urls)
    fmt.Print(string(bytes))
}
