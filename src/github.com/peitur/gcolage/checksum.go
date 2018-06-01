package gcolage

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"log"
	"os"
)

type FileChecksum struct {
	Filename  string `json:"filename"`
	Algorithm string `json:"algorithm"`
	Checksum  string `json:"checksum"`
}

func SelectHashFunc(alg string) hash.Hash {
	if alg == "sha1" {
		return sha1.New()
	} else if alg == "sha256" {
		return sha256.New()
	} else if alg == "sha512" {
		return sha512.New()
	} else {
		return SelectHashFunc("sha256")
	}
}

func HashDataFile(filename string, alg string) FileChecksum {
	fd, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	var hsh = SelectHashFunc(alg)
	if _, err := io.Copy(hsh, fd); err != nil {
		log.Fatal(err)
	}

	var chs FileChecksum
	chs.Filename = filename
	chs.Algorithm = alg
	chs.Checksum = hex.EncodeToString(hsh.Sum(nil))

	return chs
}
