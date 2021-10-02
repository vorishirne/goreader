package reader

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var ErrorNoFileName = fmt.Errorf("no file to write")

// there are already two functions that give control over length of data read
// io.LimitReader() to limit max bytes of a reader. It will put reader's EOF near
// io.ReadFull() to read to a max amount of bytes.

// there are also ones for using a delimiter
// bufio.reader.ReadByte()
// bufio.reader.ReadString()

// however, a combination of both use cases needs to be implemented via
// bufio.reader.ReadSlice() and its amazing.

func ReadSliceLimit(b *bufio.Reader, maxBytes int, delim byte) ([]byte, error) {
	b = bufio.NewReaderSize(b, maxBytes)
	return b.ReadSlice(delim)
}

// there needs to be a function that can call a callback on each line of a file,
// + should also collect error responses for each line in a map.

func CallbackOnEachLine(filePath string, callback func(string) error) (errMap *map[string]string, err error) {
	errMap = &map[string]string{}
	filePointer, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer filePointer.Close()
	fileReader := bufio.NewReader(filePointer)
	for {
		line, err := fileReader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				err := callback(string(line))
				if err != nil {
					(*errMap)[string(line)] = err.Error()
				}
				break
			}
			return nil, err
		}
		err = callback(string(line))
		if err != nil {
			(*errMap)[string(line)] = err.Error()
		}
	}
	return errMap, nil
}

func GetDirPathAndFileName(filePath string, isFilePathNonRoot bool) (dirPath string, fileName string, err error) {
	if isFilePathNonRoot {
		filePath = strings.Trim(filePath, "/")
	} else if filePath != "/" {
		filePath = strings.TrimSuffix(filePath, "/")
	}
	lastSlash := strings.LastIndexAny(filePath, "/")
	fileName = filePath[lastSlash+1:]
	if fileName == "" {
		err = ErrorNoFileName
	}
	if lastSlash == -1 {
		err = fmt.Errorf("no directory to write file to")
		return
	}
	dirPath = dirPath[:lastSlash]

	return
}
