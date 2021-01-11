package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("You need to specify a path.")
		return
	}

	root, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println("Cannot get absolute path:", err)
		return
	}

	fmt.Println("Listing files in", root)

	var c struct {
		files int
		dirs  int
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		// walk the tree to count files and folders

		if info.IsDir() {
			c.dirs++
		} else {
			c.files++
		}

		fmt.Println("_", path)
		return nil
	})

	fmt.Printf("Total: %d files in %d directories", c.files, c.dirs)
}
