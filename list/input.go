package list

type FileTree struct {
	InputPaths      *[]string
	IsDir        bool
	HasFile      bool
	MaxNestedDir int
	IsNestedDir     bool
	FilePaths       []string
	DirPaths        []string
	HiddenFilePaths []string
	HiddenDirPaths  []string
	StdinValid      bool
	SkipHiddenDirs  bool
	IsDirAllowed  func(path string) bool
	IsFileAllowed func(path string) bool
}

// maxRecursion: -1 for no limit

func Read(path *[]string, maxRecursion int) *FileTree {
	return &FileTree{InputPaths: path, MaxNestedDir: maxRecursion,
		IsDirAllowed: returnTrue,IsFileAllowed: returnTrue}
}

func returnTrue(_ string) bool {
	return true
}