package main

import (
	"github.com/wbrush/random-file/os_services"
	"fmt"
	"github.com/wbrush/random-file/process"
	"flag"
)

func main() {

	pathPtr := flag.String("path", ".", "directory with files")
	writePtr := flag.Bool("write", false, "write files")

	flag.Parse()

	path := *pathPtr
	wflag := *writePtr

	//fmt.Println("path: ", path)
	//fmt.Println("write: ", *writePtr)

	//  read files in directory
	//path := "c:\\Projects\\golang\\src\\common\\randomfile\\test"
	origList, err := os_services.ReadDirectory(path)
	if (err != nil) {
		fmt.Println("Error Reading Files: ", err)
		return
	}

	//  strip off preceding number if it is there
	fileList, err := process.StripNames(origList)
	if (err != nil) {
		fmt.Println("Error In File Names: ", err)
		return
	}
	//  generate random number list
	randomNumList := process.GenerateRandomNumbers(int(len(fileList)))

	//  create new file names
	finalList, err := process.CreateNames (fileList, randomNumList)

	//  write new names
	//fmt.Println("Number of Files: ", len(finalList))
	//for _,item := range randomNumList {
	//	fmt.Println("item:",item)
	//}
	//for i,fname := range finalList {
	//	fmt.Println("original ", origList[i], " becomes fname:",fname)
	//}

	os_services.RenameFiles(path, origList, fileList, finalList, wflag)

	return
}
