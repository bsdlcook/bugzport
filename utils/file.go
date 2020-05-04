package utils

import (
	"io"
	"os"
)

func CopyFile(source string, destination string) {
	sourceFile, _ := os.Open(source)
	defer sourceFile.Close()

	newFile, _ := os.Create(destination)
	defer newFile.Close()

	io.Copy(newFile, sourceFile)
}
