package gcolage

import (
	"fmt"
	"log"
	"os"
)

func CreateTempDir(iroot, iprefix string) (string, error) {
	var path string = fmt.Sprintf("%s/%s_%s", iroot, iprefix, RandomString(10))

	crtd := os.MkdirAll(path, 0700)
	if crtd != nil {
		log.Printf("ERROR: Could not create temporary directory %s : %s\n", path, crtd)
		return "", crtd
	}

	return path, nil
}

func RemoveDirTree(path string) error {
	return os.RemoveAll(path)
}

func CopyFile(src string, trgdir string) error {
	return nil
}

func UnpackZip(filename string, trgdir string) (string, error) {
	return "", nil
}

func UnpackTarGz(filename string, trgdir string) (string, error) {
	return "", nil
}
