package list

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func (tree *FileTree) UpdateFiles() error {
	nestedLevel := 0
	//nestedPathTree := &NestedPathTree{ }
	for _, inputPath := range *tree.InputPaths {

		inputPath = strings.TrimSpace(inputPath)
		file, err := os.Stat(inputPath)
		if err != nil {
			return err
		}
		if file.IsDir() && tree.MaxNestedDir != nestedLevel &&
			tree.IsDirAllowed(inputPath) {
			var nTree = tree.GeneratePathTree(inputPath)
			if tree.GenerateNestedTree {
				tree.NestedPathTree[inputPath] = nTree
			}
			tree.IsDir = true
			nestedLevel_ := nestedLevel + 1
			// . actually means all the dirs and files in curr directory
			// not the name of a dir.
			// it's just a shorthand. iterateDir will just expand this shorthand,
			// and must not be treated as one nested level of dir
			if inputPath == "." || inputPath == "./" {
				nestedLevel_ = nestedLevel
			}
			err = tree.iterateDir(&inputPath, nestedLevel_, nTree)
			if err != nil {
				return err
			}
		} else if tree.IsFileAllowed(inputPath) {
			if tree.GenerateNestedTree {
				tree.NestedPathTree[inputPath] = nil
			}
			tree.HasFile = true
			if CheckPathHidden(inputPath) {
				tree.HiddenFilePaths = append(tree.HiddenFilePaths, inputPath)
			} else {
				tree.FilePaths = append(tree.FilePaths, inputPath)
			}
		}
	}
	return nil
}

func (tree *FileTree) iterateDir(basePath *string, nestedLevel int, nTree *NestedPathTree) error {

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
		if file.IsDir() && tree.MaxNestedDir != nestedLevel &&
			tree.IsDirAllowed(path) {
			tree.AddDirToPathTree(nTree, file.Name(), *basePath)
			tree.IsNestedDir = true

			err := tree.iterateDir(&path, nestedLevel+1, nTree.NestedDir[file.Name()])
			if err != nil {
				return err
			}
		} else if tree.IsFileAllowed(path) {
			tree.HasFile = true
			tree.AddFileToPathTree(nTree, file.Name(), *basePath)
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

func (tree *FileTree) AddDirToPathTree(nTree *NestedPathTree, dir string, basePath string) {
	if tree.GenerateNestedTree {
		npt := &NestedPathTree{
			FullPathTillNow: path.Join(basePath, dir),
			IsHiddenDir:     CheckPathHidden(path.Join(basePath, dir)),
			NestedDir:       make(map[string]*NestedPathTree),
		}
		nTree.NestedDir[dir] = npt
		if tree.GenerateNestedTreeFileOrder {
			nTree.FileOrder = append(nTree.FileOrder, fileOrder{npt.FullPathTillNow, npt})
		}
	}
}

func (tree *FileTree) AddFileToPathTree(nTree *NestedPathTree, file string, basePath string) {
	if tree.GenerateNestedTree {
		nTree.FilesHere = append(nTree.FilesHere, path.Join(basePath, file))
		if tree.GenerateNestedTreeFileOrder {
			nTree.FileOrder = append(nTree.FileOrder, fileOrder{path.Join(basePath, file), nil})
		}
	}
}

func (tree *FileTree) GeneratePathTree(basePath string) *NestedPathTree {
	return &NestedPathTree{
		FullPathTillNow: basePath,
		IsHiddenDir:     CheckPathHidden(basePath),
		NestedDir:       make(map[string]*NestedPathTree)}
}
