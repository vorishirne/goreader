package list

type FileTree struct {
	InputPaths       []string
	IsDir            bool
	IsFile           bool
	NestedDirAllowed bool
	IsNestedDir      bool
	FilePaths        []string
	DirPaths         []string
	StdinValid       bool
}

func Read(path []string) *FileTree {
	return &FileTree{InputPaths: path}
}
