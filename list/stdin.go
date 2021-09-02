package list

import (
	"os"
)

//check if stdin is good input source

func (tree *FileTree) UpdateStdinStatus() {
	// just try to access the stdin, no matter if its has something or not, it is always accessible.
	// input to stdin gets into an ephemeral file, and we are fetching stats to that file
	stdinFileInfo, err := os.Stdin.Stat()
	if err != nil {
		return
	}
	// stdin must have mode of char device(that is tty).
	// if the mode is not what is bitwise "&" with, it will not return 0.
	// (Mode has info for file, in a bitwise operatable manner)

	// this will be false when
	// app
	// app > something
	// app arg
	// app | grep

	// this will be true when
	// echo "sdf" | app
	// app <<< "sdf"

	if stdinFileInfo.Mode()&os.ModeCharDevice == 0 {
		//fmt.Println("TTY")
	} else {
		return
	}

	//// returns zero in case of pipe
	// fmt.Println("stdin size: ",stdinFileInfo.Size())

	//// how to read
	//// here size belongs to the size of inbuilt buffer. i.e. 4096 bytes/4 kb sometimes
	//textReader := bufio.NewReader(os.Stdin)
	//bytes,_,_:=  textReader.ReadLine()
	//bytes2,_,_:=  textReader.ReadLine()
	//bytes3,_,_:=  textReader.ReadLine()
	//fmt.Println("input bytes: ",string(bytes))
	//fmt.Println("input bytes: ",string(bytes2))
	//fmt.Println("input bytes: ",string(bytes3))

	tree.StdinValid = true
}
