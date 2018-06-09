package gcolage

import (
	"regexp"
)

func ConcatBytes(a []byte, b []byte) []byte {
	var res []byte = a
	for _, x := range b {
		res = append(res, x)
	}
	return res
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func ApplyVersionString(v string, str string) string {
	return regexp.MustCompile("<%\\s*version\\s*%>").ReplaceAllString(str, v)
}

func ApplyProductString(v string, str string) string {
	return regexp.MustCompile("<%\\s*product\\s*%>").ReplaceAllString(str, v)
}
