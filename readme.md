# goreader
A utility library to do files/io/bytes processing/parsing in file-system or network. 

1. These features are really common to be implemented for a non-server/local-running application. Though looking fairly easy, almost all had edge cases to be taken care of minutely.
2. To use a library package instead of custom segments saves redundant testing/debugging.

# Feature list 
* generate parse tree for given paths
  * list out files/folders separately
  * Get hidden files/folders separately
  * Preserve nested order for dir and files
  * Control recursion level
  * Pass custom functions to filter out files and dirs
  * Include StdIn if valid
* Encodings
  * json to 
    * file
    * yaml
  * yaml to
    * json
    * file
  * file to
    * json 
    * yaml
* IO ops
  * copy multi files to single one
  * clone file
  * Pass a callback to be executed over each line of a file
  * bash `dirname` & `basedir` function
# How to use

#### use lister
```golang

tree := list.Read(&[]string{"."}, 2)
//tree.SkipHiddenDirs = true
e := tree.UpdateFiles()
tree.UpdateStdinStatus()

if e != nil {
    log.Fatal(e)
}
log.Println(tree.IsNestedDir)
```

#### use reader
```golang
for _,path := range tree.FilePaths{
	if strings.HasSuffix(path,"html") || strings.HasSuffix(path,"htm"){
		err := reader.CloneFile(path, path+".bak")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = reader.CopyContent(path, "envoy/css.ht", path+".bak")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Remove(path+".bak")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
```