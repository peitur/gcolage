package gcolage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
)

type FileCollectorSpec struct {
	Version   string `json:"version"`
	Url       string `json:"url"`
	Filename  string `json:"filename"`
	Signature string `json:"signature"`
}

type FileCollectorSpecs struct {
	Specs []FileCollectorSpec
}

func LoadCollectorConfigFiles(dir string) []FileCollectorSpecs {
	var data []FileCollectorSpecs

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		m, _ := regexp.MatchString(".*.json$", f.Name())
		if m == true {
			path := dir + "/" + f.Name()
			log.Printf("Loading config from: %s", path)

			sps := FileCollectorSpecs{}
			spl := LoadCollectorConfigSpecs(path)
			sps.Specs = spl
			data = append(data, sps)
		}
	}

	return data
}

func SaveCollectorFileInfo(filename string, target_dir string, info FileInfo) {
	var fi, _ = json.MarshalIndent(info, "", " ")
	var path = target_dir + "/" + filename
	ioutil.WriteFile(path, fi, 0644)
}

func LoadCollectorConfigSpecs(filename string) []FileCollectorSpec {
	buffer, err := ReadFileRaw(filename)
	if err != nil {
		log.Fatal(err)
	}

	//	var conf FileCollectorSpecs

	spec := []FileCollectorSpec{}
	err = json.Unmarshal([]byte(buffer), &spec)
	if err != nil {
		log.Fatal(err)
	}

	return spec
}
