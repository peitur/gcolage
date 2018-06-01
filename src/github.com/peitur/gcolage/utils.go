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

	re := regexp.MustCompile("<%\\s*version\\s*%>")
	return re.ReplaceAllString(str, v)
}
