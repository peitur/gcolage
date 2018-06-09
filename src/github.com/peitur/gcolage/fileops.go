package gcolage

import (
	"bufio"
	"log"
	"os"
	"syscall"
	"time"
)

type FileInfo struct {
	Filename string       `json:"filenmae"`
	Size     int64        `json:"size"`
	Ctime    time.Time    `json:"ctime"`
	Mtime    time.Time    `json:"mtime"`
	Atime    time.Time    `json:"atime"`
	Checksum FileChecksum `json:"checksum"`
}

type FileInfoList struct {
	Files []FileInfo
}

type FileInfoTracker struct {
	File   FileInfo `json:"file"`
	Source string   `json:"source"`
}

func GetFileInfo(filename string) (FileInfo, error) {
	var stt, err = os.Stat(filename)
	var t FileInfo

	if err == nil {

		t.Filename = filename
		t.Size = stt.Size()
		t.Mtime = stt.ModTime()

		//  	var stat = stt.Sys()
		var stat = stt.Sys().(*syscall.Stat_t)
		t.Atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
		t.Ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))

		t.Checksum = HashDataFile(filename, "sha256 ")

		return t, nil
	}

	return t, err
}

func ReadFileRaw(filename string) ([]byte, error) {
	var res []byte
	fd, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	scr := bufio.NewScanner(fd)
	for scr.Scan() {
		res = ConcatBytes(res, scr.Bytes())
	}

	if err := scr.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}

func FileExists(filename string) bool {
	return false
}
