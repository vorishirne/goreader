package reader

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
)

func CopyContent(toPath string, fromPaths ...string) error {
	toDir := filepath.Dir(toPath)
	if _, err := os.Stat(toDir); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(toDir, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

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
		ru, nRune, err := fileReader.ReadRune()
		if err != nil || nRune < 1 {
			return err
		}
		if ru != '\ufeff' {
			err := fileReader.UnreadRune()
			if err != nil {
				return err
			}
		}

		nBytes, err := fileReader.WriteTo(fileToWrite)
		if err2 := fileToRead.Close(); err2 != nil {
			return err2
		}
		if err != nil || nBytes < 1 {
			return err
		}

		_, err = fileToWrite.WriteString("\n")
		if err != nil {
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
