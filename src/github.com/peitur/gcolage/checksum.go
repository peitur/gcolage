package gcolage

import (
  "hash"
  "crypto/sha1"
  "crypto/sha256"
  "crypto/sha512"
  "encoding/hex"
  "os"
  "io"
  "log"
)

type FileChecksum struct{
	filename string
	algorithm string
	checksum string
}

func SelectHashFunc( alg string ) ( hash.Hash ){
  if alg == "sha1"{
    return sha1.New()
  }else if alg == "sha256"{
    return sha256.New()
  }else if alg == "sha512"{
    return sha512.New()
  }else{
    return SelectHashFunc( "sha256" )
  }
}

func HashDataFile( filename string, alg string ) ( FileChecksum ){
  fd, err := os.Open( filename )
  if err != nil {
    log.Fatal(err)
  }
  defer fd.Close()

  var hsh = SelectHashFunc( alg )
  if _, err := io.Copy( hsh, fd ); err != nil {
    log.Fatal(err)
  }

  var chs FileChecksum
  chs.filename = filename
  chs.algorithm = alg
  chs.checksum = hex.EncodeToString( hsh.Sum(nil))

  return chs
}
