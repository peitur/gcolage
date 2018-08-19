package gcolage

import (
	"encoding/json"
	"fmt"
	"log"
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

func JsonGenericSlice(buffer []byte) ([]interface{}, error) {
	var gen interface{}
	err := json.Unmarshal(buffer, &gen)
	if err != nil {
		return nil, err
	}

	return gen.([]interface{}), nil
}

func JsonTest(filename string) {
	buffer, err := ReadFileRaw(filename)
	if err != nil {
		log.Fatal(err)
	}

	js, _ := JsonGenericSlice(buffer)
	for i := range js {

		f := js[i].(map[string]interface{})
		for k := range f {
			fmt.Println(k)
		}
	}
}
