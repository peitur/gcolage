package gcolage

import (
	"encoding/json"
	"bufio"
	"log"
	"os"
)


func ConcatBytes( a []byte, b []byte ) []byte {
	var res []byte = a
	for _, x := range b {
		res = append( res, x )
	}
	return res
}

func ReadTextFile( filename string ) []byte {
	var res []byte
	fd, err := os.Open( filename )
	if err != nil {
	    log.Fatal(err)
	}
	defer fd.Close()

	scr := bufio.NewScanner( fd )
	for scr.Scan() {
	  res = ConcatBytes( res, scr.Bytes() )
	}

	if err := scr.Err(); err != nil {
	    log.Fatal(err)
	}

	return res
}


func BytesToString(data []byte) string {
	return string(data[:])
}

func JsonEncodeConfig( data []byte ) {

}

func JsonDecodeConfig( data []byte ) ( map[string]*json.RawMessage ){
	var objmap map[string]*json.RawMessage
	json.Unmarshal( data, &objmap )
	return objmap
}
