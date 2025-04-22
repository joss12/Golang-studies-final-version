package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// tempFileName := "temporaryFile"

	// tempFile, err := os.CreateTemp("", tempFileName)
	// checkError(err)

	// fmt.Println("Temporary file created:", tempFile.Name())

	// defer os.Remove(tempFile.Name())
	// defer tempFile.Close()

	//Create temporary directories
	tempDir, err := os.MkdirTemp("", "GoCourseTempDir")
	checkError(err)

	// defer os.Remove(tempDir)
	defer os.RemoveAll(tempDir)

	fmt.Println("Temporary Directory created:", tempDir)

}
