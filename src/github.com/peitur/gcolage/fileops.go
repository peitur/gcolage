package gcolage

import (
	"time"
	"os"
	"syscall"
  "io/ioutil"
)


type FileInfo struct {
	filename string
	size int64
	ctime time.Time
	mtime time.Time
	atime time.Time
	checksum FileChecksum
}

func GetFileInfo( filename string ) ( FileInfo, error) {
	var stt, err = os.Stat( filename )
	var t FileInfo

	if err == nil{

		t.filename = filename
    t.size = stt.Size()
		t.mtime = stt.ModTime()

		var stat = stt.Sys().(*syscall.Stat_t)
		t.atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
		t.ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))

		return t, nil
	}

	return t, err
}

func ReadFileRaw( filename string ) ( []byte, error ){
  return ioutil.ReadFile( filename )
}

func WriteFileRaw( filename string, data []byte ) ( uint, error ){
  return 0, nil
}
