package gcolage

import (
  "encoding/json"
  "log"
)

type FileCollectorConfig struct {
  ConfigPath string
  TargetPath string
  Checksum string
}

type FileCollectorSpec struct {
  Version string
  Url string
  Filename string
  Signature string
}

type FileCollectorSpecs struct {
  Collection []FileCollectorSpec
}

type PipCollectorConfig struct {
  ConfigPath string
  TargetPath string
  Checksum string
}



func ReadConfigFile( filename string ) FileCollectorConfig {
  var buffer []byte = ReadTextFile( filename )
  var conf FileCollectorConfig
	err := json.Unmarshal( []byte( buffer ), &conf )
  if err != nil {
    log.Fatal( err )
  }
  return conf
}
