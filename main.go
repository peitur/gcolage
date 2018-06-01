package main

import (
	"fmt"
	"log"

	"github.com/peitur/gcolage"
)

func main() {
	fmt.Println("Welcome ...")

	var filename = "src/github.com/peitur/gcolage/test/sample.json"
	var configfile = "test/config.json"

	var t, e = gcolage.GetFileInfo(filename)
	if e != nil {
		log.Fatal(e)
	}

	log.Println(">>> ", t)

	fmt.Println("----------------------------------------------------------")
	conf := gcolage.ReadConfigFile(configfile)
	fmt.Println(gcolage.LoadCollectorConfigFiles(conf.ConfigPath))
	//fmt.Println(gcolage.GetFileInfo(filename))
	fmt.Println("----------------------------------------------------------")
	fmt.Println(gcolage.ApplyVersionString("1.10.2", "https://dl.google.com/go/go<% version %>.linux-amd64.tar.gz-<% version %>"))
	fmt.Println("----------------------------------------------------------")
	fi, _ := gcolage.GetFileInfo(filename)
	gcolage.SaveCollectorFileInfo("x.json", "/tmp", fi)
	fmt.Println("----------------------------------------------------------")

	fmt.Println("----------------------------------------------------------")

}
