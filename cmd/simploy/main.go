package main

import (
	"fmt"
	"os"
	"path/filepath"
	"simploy/pkg/fs"
)

func printCurrentDir() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println("Current Dir: " + exPath)
}

func main() {

	printCurrentDir()

	os.Chmod("test", 0777)

	src := "test"
	dst := "test2"

	println("Copying", src, "to", dst)

	err := fs.CopyDir(src, dst)

	if err != nil {
		println("Error copying", src, "to", dst)
		println(err)
	}
}
