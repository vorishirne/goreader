package list

type fileOrder struct {
	Path           string
	NestedPathTree *NestedPathTree
}
type NestedPathTree struct {
	NestedDir       map[string]*NestedPathTree
	FileOrder       []fileOrder
	FullPathTillNow string
	FilesHere       []string
	IsHiddenDir     bool
}
type FileTree struct {
	InputPaths                  *[]string
	IsDir                       bool
	HasFile                     bool
	MaxNestedDir                int
	IsNestedDir                 bool
	FilePaths                   []string
	DirPaths                    []string
	HiddenFilePaths             []string
	HiddenDirPaths              []string
	StdinValid                  bool
	SkipHiddenDirs              bool
	GenerateNestedTree          bool
	GenerateNestedTreeFileOrder bool
	NestedPathTree              map[string]*NestedPathTree
	IsDirAllowed                func(path string) bool
	IsFileAllowed               func(path string) bool
}

// maxRecursion: -1 for no limit

func Read(path *[]string, maxRecursion int) *FileTree {
	return &FileTree{InputPaths: path, MaxNestedDir: maxRecursion,
		IsDirAllowed: returnTrue, IsFileAllowed: returnTrue,
		NestedPathTree: make(map[string]*NestedPathTree),
	}
}

func returnTrue(_ string) bool {
	return true
}
