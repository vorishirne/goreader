package list

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func (tree *FileTree) UpdateFiles() error {

	for _, path := range tree.InputPaths {
		fInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if fInfo.IsDir() {
			tree.IsDir = true
			tree.DirPaths = append(tree.DirPaths, path)
			err = tree.iterateDir(&path)
			if err != nil {
				return err
			}
		} else {
			tree.IsFile = true
			tree.FilePaths = append(tree.FilePaths, path)
		}
	}
	return nil
}

func (tree *FileTree) iterateDir(basePath *string) error {
	ls, err := ioutil.ReadDir(*basePath)
	if err != nil {
		return err
	}
	for _, file := range ls {
		if file.IsDir() && tree.NestedDirAllowed {
			path := filepath.Join(*basePath, file.Name())
			tree.DirPaths = append(tree.DirPaths, path)
			tree.IsNestedDir = true

			err := tree.iterateDir(&path)
			if err != nil {
				return err
			}

		} else {
			tree.IsFile = true
			tree.FilePaths = append(tree.FilePaths, filepath.Join(*basePath, file.Name()))
			// now use
			// io.LimitReader() to limit max bytes of a reader. It will put reader's EOF near
			// io.ReadFull() to read to a max amount of bytes.
		}
	}
	return nil
}
