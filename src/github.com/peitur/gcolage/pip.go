package gcolage

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PipProjectInfoJson(project string) string {
	return fmt.Sprintf("https://pypi.org/pypi/%s/json", project)
}

/*
Get project info
* name

*/
func PipRequestProjectInfo(project string) (map[string]interface{}, error) {
	resp, err := http.Get(PipProjectInfoJson(project))
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	d, err := JsonGenericMap(body)
	if err != nil {
		return nil, err
	}

	return d, nil
}

/*
Find requirements by
1. download package
2. unpack package to temp dir
3. read requirements files (requirements.txt and ... )
4. check
*/
func PipReadRequirements(filename string) []string {
	var res []string
	buffer, err := ReadFileRaw(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buffer))
	return res
}
