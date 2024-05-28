package main

import "os"
import "encoding/json"

// loadBookworms reads the file and return the list of bookworms, and their beloved books, found therein.
func loadBookworms(filePath string) ([]Bookworm, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil
	}
	defer f.Close()

	// Initialize the type in which the file will be decoded.
	var bookworms []Bookworm

	// Decode the file and store the content in the variable bookworms.
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil

}

// Bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name string  `json:"name"`
	Book []Books `json:"books"`
}

// Books describes a book on a bookworm's shelf
type Books struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}
