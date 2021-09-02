package list

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func (tree *FileTree) UpdateFiles() error {
	nestedLevel := 0
	for _, inputPath := range *tree.InputPaths {
		inputPath = strings.TrimSpace(inputPath)
		file, err := os.Stat(inputPath)
		if err != nil {
			return err
		}
		path := filepath.Join(inputPath, file.Name())
		if file.IsDir() && tree.MaxNestedDir != nestedLevel {
			tree.IsDir = true
			nestedLevel_ := nestedLevel + 1
			// . actually means all the dirs and files in curr directory
			// not the name of a dir.
			// it's just a shorthand. iterateDir will just expand this shorthand,
			// and must not be treated as one nested level of dir
			if path == "." || path == "./" {
				nestedLevel_ = nestedLevel
			}
			err = tree.iterateDir(&inputPath, nestedLevel_)
			if err != nil {
				return err
			}
		} else {
			tree.IsFile = true
			if CheckPathHidden(path) {
				tree.HiddenFilePaths = append(tree.HiddenFilePaths, path)
			} else {
				tree.FilePaths = append(tree.FilePaths, path)
			}
		}
	}
	return nil
}

func (tree *FileTree) iterateDir(basePath *string, nestedLevel int) error {

	if CheckPathHidden(*basePath) {
		if tree.SkipHiddenDirs {
			return nil
		}
		tree.HiddenDirPaths = append(tree.HiddenDirPaths, *basePath)
	} else {
		tree.DirPaths = append(tree.DirPaths, *basePath)
	}
	ls, err := ioutil.ReadDir(*basePath)
	if err != nil {
		return err
	}
	for _, file := range ls {
		path := filepath.Join(*basePath, file.Name())
		if file.IsDir() && tree.MaxNestedDir != nestedLevel {
			tree.IsNestedDir = true

			err := tree.iterateDir(&path, nestedLevel+1)
			if err != nil {
				return err
			}
		} else {
			tree.IsFile = true

			if CheckPathHidden(path) {
				tree.HiddenFilePaths = append(tree.HiddenFilePaths, path)
			} else {
				tree.FilePaths = append(tree.FilePaths, path)
			}
			// now use
			// io.LimitReader() to limit max bytes of a reader. It will put reader's EOF near
			// io.ReadFull() to read to a max amount of bytes.
		}
	}
	return nil
}

func CheckPathHidden(path string) bool {
	//todo(velcrine): replace this with a regex. Be pro
	if path == "." {
		return false
	}
	if strings.HasPrefix(path, ".") && !strings.HasPrefix(path, "./") {
		return true
	}
	if strings.Contains(path, "/.") {
		return true
	}
	return false
}
