package gcolage

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PipInfoData struct {
	Version string
}

type PipReleaseData struct {
	Checksum       string
	Filename       string
	PythonVersion  string
	PythonRequered string
	PackageType    string
	UploadTime     string
	Version        string
	Size           uint64
	Url            string
	//	Request        map[string]interface{}
}

type PipUrlData PipReleaseData

type PipProjectData struct {
	Info     PipInfoData
	Releases []PipReleaseData
	Urls     []PipReleaseData
}

func PipProjectInfoJson(project string) string {
	return fmt.Sprintf("https://pypi.org/pypi/%s/json", project)
}

func PipParseReleaseInfoData(ver string, relx map[string]interface{}) PipReleaseData {
	var r PipReleaseData

	r.Version = ver
	r.Filename, _ = relx["filename"].(string)
	r.PackageType, _ = relx["packagetype"].(string)
	r.PythonVersion, _ = relx["python_version"].(string)
	r.PythonRequered, _ = relx["required_python"].(string)
	r.UploadTime, _ = relx["upload_time"].(string)
	r.Size, _ = relx["size"].(uint64)
	r.Url, _ = relx["url"].(string)
	//	r.Request = relx

	return r
}

/*
Get project info
* name

*/
func PipCurrentReleaseInfo(ver string, urls []interface{}) []PipReleaseData {
	var res []PipReleaseData
	for u := range urls {
		res = append(res, PipParseReleaseInfoData(ver, urls[u].(map[string]interface{})))
	}
	return res
}

func PipAllReleaseInfo(releases map[string]interface{}) []PipReleaseData {
	var rrels []PipReleaseData

	for ver := range releases {
		rel := releases[ver].([]interface{})
		for j := range rel {
			rrels = append(rrels, PipParseReleaseInfoData(ver, rel[j].(map[string]interface{})))
		}
	}

	return rrels
}

func PipRequestProjectInfo(project string) (PipProjectData, error) {
	log.Printf("INFO: Getting info for %s \n", project)

	var res PipProjectData

	resp, err := http.Get(PipProjectInfoJson(project))
	if err != nil {
		log.Printf("ERROR getting project %s : %s", project, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ERROR: Reading body failed: %s\n", err)
		return res, err
	}

	d, err := JsonGenericMap(body)
	if err != nil {
		log.Printf("ERROR: Decoding reply failed: %s\n", err)
		return res, err
	}

	info := d["info"].(map[string]interface{})
	releases := d["releases"].(map[string]interface{})
	urls := d["urls"].([]interface{})

	var rinf PipInfoData

	rinf.Version, _ = info["version"].(string)

	res.Info = rinf
	res.Releases = PipAllReleaseInfo(releases)
	res.Urls = PipCurrentReleaseInfo(rinf.Version, urls)
	return res, nil
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

	fmt.Println(buffer)
	return res
}
