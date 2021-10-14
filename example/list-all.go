package main

import (
	"github.com/watergist/file-engine/list"
	"log"
)

func main() {
	tree := list.Read(&[]string{"../goreader/blog"}, 2)
	tree.GenerateNestedTree = true
	//tree.SkipHiddenDirs = true
	e := tree.UpdateFiles()
	tree.UpdateStdinStatus()

	if e != nil {
		log.Fatal(e)
	}
	log.Println(tree.IsNestedDir)

	//// to try the reader package
	//for _,path := range tree.FilePaths{
	//	if strings.HasSuffix(path,"html") || strings.HasSuffix(path,"htm"){
	//		err := reader.CloneFile(path, path+".bak")
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		err = reader.CopyContent(path, "envoy/css.ht", path+".bak")
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		err = os.Remove(path+".bak")
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//	}
	//}
}
