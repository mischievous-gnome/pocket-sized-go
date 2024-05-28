package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json") // will be completed along the way
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load booworms: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookworms)
}
