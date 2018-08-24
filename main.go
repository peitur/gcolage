package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/peitur/gcolage"
)

func main() {
	fmt.Println("Welcome ...")

	var configfile = "test/config.json"

	var t, e = gcolage.GetFileInfo(configfile)
	if e != nil {
		log.Fatal(e)
	}

	var mode string = os.Args[1]
	log.Printf("Mode: %s\n", mode)
	log.Printf("Filename: %s\n", configfile)
	log.Printf(">> Size: %d\n", t.Size)
	log.Printf(">> Checksum: %s: %s\n", t.Checksum.Algorithm, t.Checksum.Checksum)

	if mode == "get" {
		config := gcolage.ReadConfigFile(configfile)
		log.Printf("Loding configfile: %s\n", configfile)
		specs := gcolage.LoadCollectorConfigFiles(config.ConfigPath)

		for spec := range specs {
			for spc := range specs[spec].Specs {
				s := specs[spec].Specs[spc]
				s_Url := gcolage.ApplyVersionString(s.Version, s.Url)
				s_Filename := s.Filename
				if len(s_Filename) == 0 {
					xs := regexp.MustCompile("/").Split(s_Url, -1)
					s_Filename = xs[len(xs)-1]
				}

				s_Filename = gcolage.ApplyVersionString(s.Version, s_Filename)
				t_Filename := config.TargetPath + "/" + s_Filename
				log.Printf(">>> [%s] %s\n", s_Filename, s_Url)
				if !gcolage.FileExists(t_Filename) {
					log.Printf("||>>> Downloading %s -> %s\n", s_Url, t_Filename)
					gcolage.RequestFile(s_Url, t_Filename)

				} else {
					log.Printf("||>>> Already downloaded %s -> skipping\n", s_Url)
				}
			}
		}
	} else if mode == "pip" {
		pkglistfile := "requirements.txt"
		fmt.Printf(">> Requirements file: %s\n", pkglistfile)

		gcolage.PipRequestProjectInfo("kiwi")

	} else {
		log.Panic("Method not supported, Supported methods are [get]")
	}
	/*
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
		fmt.Println(gcolage.RequestFile("https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz", "/tmp/stuff.xml"))
		fmt.Println("----------------------------------------------------------")
	*/
}
