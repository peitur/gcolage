package gcolage

import "fmt"

func CreateTempDir(iroot, iprefix string) (string, error) {
	var path string = fmt.Sprintf("%s/%s_%s", iroot, iprefix, RandomString(10))

	return path, nil
}

func RemoveDirTree(path string) error {
	return nil
}
