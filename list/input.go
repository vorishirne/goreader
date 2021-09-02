package list

type FileTree struct {
	InputPaths      *[]string
	IsDir           bool
	IsFile          bool
	MaxNestedDir    int
	IsNestedDir     bool
	FilePaths       []string
	DirPaths        []string
	HiddenFilePaths []string
	HiddenDirPaths  []string
	StdinValid      bool
	SkipHiddenDirs  bool
}

// maxRecursion: -1 for no limit

func Read(path *[]string, maxRecursion int) *FileTree {
	return &FileTree{InputPaths: path, MaxNestedDir: maxRecursion}
}
