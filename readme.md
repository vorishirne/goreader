# file-engine
module to deal with large or more files. 
#### side-note: this module will not require frequent maintenance as it is quite close to the golang std library.

file engine has two utilities:
    
1. list: takes away the headache of 
    dealing with directories and files, hidden ones, nested ones.
       
    It's kind of "ls" command as a package, but doesn't intends
    to show file properties, permissions or ownership.
       
    General packages just take care of reading and accessing
    files instead of dealing with file permissions.


2. reader: binds frequently used file io functions 

    for ex. cloning a file, add to front of a file(reverse append), etc.

    These functions are frequent to use, but not provided by standard go io package.

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