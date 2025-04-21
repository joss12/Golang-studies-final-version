package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		// fmt.Println(err)
	}
}

func main() {
	// **Create directories
	// err := os.Mkdir("subdir", 0755)
	// checkError(err)
	// checkError(os.Mkdir("subdir1", 0755))

	// **Delete directories
	// defer os.RemoveAll("subdir1")

	// **create files
	// os.WriteFile("subdir1/file", []byte(""), 0755)

	// *Create multiple directories
	// checkError(os.MkdirAll("subdir/parent/child1", 0755))
	// checkError(os.MkdirAll("subdir/parent/child2", 0755))
	// checkError(os.MkdirAll("subdir/parent/child3", 0755))

	// os.WriteFile("subdir/parent/file", []byte(""), 0755)
	// os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	// **Read directories
	result, err := os.ReadDir("subdir/parent")
	checkError(err)
	for _, v := range result {
		fmt.Println(v.Name(), v.IsDir(), v.Type())
	}

	checkError(os.Chdir("subdir/parent/child"))
	checkError(os.Chdir("../../.."))
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)
	result, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading subdir/parent/child")
	for _, v := range result {
		fmt.Println(v)
	}

	checkError(os.Chdir("../../.."))
	dir, err = os.Getwd()
	checkError(err)
	fmt.Println(dir)

	// Filepath.Walk and filepath.WalkDir
	pathfile := "subdir"
	fmt.Println("Walking directory")
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})

	checkError(err)

	//Remove all directories
	// checkError(os.RemoveAll("subdir"))
	// checkError(os.Remove("subdir1"))

}
