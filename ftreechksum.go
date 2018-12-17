package main

import (
	"fmt"
	"os"

	"github.com/peitur/gcolage"

	"github.com/pborman/getopt"
)

func main() {
	helpFlag := getopt.Bool('h', "Show help")
	pathFlag := getopt.StringLong("path", 'p', ".", "Root path")
	chksFlag := getopt.StringLong("checksum", 'c', "sha256", "Checksum to use")
	//	outpFlag := getopt.StringLong("out", 'o', "", "Output file")

	opts := getopt.CommandLine
	opts.Parse(os.Args)
	//	getopt.Getopt(nil)

	if *helpFlag == true {
		getopt.Usage()
		os.Exit(1)
	}

	files, _ := gcolage.DirectoryTree(*pathFlag)
	for p := range files {
		chksm := gcolage.HashDataFile(files[p], *chksFlag)
		fmt.Printf("%s \t %s\n", chksm.Filename, chksm.Checksum)
	}
}
