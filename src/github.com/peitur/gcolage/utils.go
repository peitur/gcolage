package gcolage

import (
	"time"
	"os"
	"syscall"
)

type FileInfo struct {
		filename string
		size int64
		ctime time.Time
		mtime time.Time
		atime time.Time
		checksum string
}

func GetFileInfo( filename string ) ( FileInfo, error) {
	var stt, err = os.Stat( filename )
	var t FileInfo

	if err == nil{

		t.filename = filename
		t.mtime = stt.ModTime()

		var stat = stt.Sys().(*syscall.Stat_t)
		t.atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
		t.ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))

		return t, nil
	}

	return t, err
}
