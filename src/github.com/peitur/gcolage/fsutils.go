package gcolage

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func DirectoryTree(rp string) ([]string, error) {
	var res []string
	var dirs []string

	files, err := ioutil.ReadDir(rp)
	if err != nil {
		return make([]string, 0), err
	}

	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, fmt.Sprintf("%s/%s", rp, file.Name()))
		} else {
			res = append(res, fmt.Sprintf("%s/%s", rp, file.Name()))
		}
	}

	for d := range dirs {
		x, _ := DirectoryTree(dirs[d])
		for f := range x {
			res = append(res, x[f])
		}
	}

	return res, nil
}

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
