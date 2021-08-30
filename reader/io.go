package reader

import (
	"bufio"
	"os"
)

func CopyContent(toPath string, fromPaths ...string) error {
	fileToWrite, err := os.OpenFile(toPath,
		os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer fileToWrite.Close()

	for _, v := range fromPaths {
		fileToRead, err := os.Open(v)
		if err != nil {
			_ = fileToRead.Close()
			return err
		}
		fileReader := bufio.NewReader(fileToRead)

		nBytes, err := fileReader.WriteTo(fileToWrite)
		if err2 := fileToRead.Close(); err2 != nil {
			return err2
		}
		if err != nil || nBytes < 1 {
			return err
		}
	}
	return nil
}

func CloneFile(from string, to string) error {
	newFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer newFile.Close()

	fileToWrite, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fileToWrite.Close()
	fileReader := bufio.NewReader(fileToWrite)
	nBytes, err := fileReader.WriteTo(newFile)
	if err != nil || nBytes < 1 {
		return err
	}
	return nil
}
