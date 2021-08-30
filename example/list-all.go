package main

import (
	"github.com/watergist/file-engine/list"
	"github.com/watergist/file-engine/reader"
	"log"
)

func main() {
	tree := list.Read([]string{"example", "list", "reader"})
	tree.NestedDirAllowed = true
	e := tree.UpdateFiles()
	if e != nil {
		log.Fatal(e)
	}
	tree.CheckStdinStatus()
	log.Println(tree.IsNestedDir)

	//
	log.Println(reader.CloneFile("go.mod", "bo.mod"))
	log.Fatal(reader.CopyContent("go.mod", "LICENSE", ".gitignore", "bo.mod"))
}
