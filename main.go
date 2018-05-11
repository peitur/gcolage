package main

import (
	"github.com/peitur/gcolage"
	"fmt"
	"log"
)


func main( ){
  fmt.Println("Welcome ...")

	var filename = "src/github.com/peitur/gcolage/test/sample.json"

  var t, e = gcolage.GetFileInfo( filename )
	if e != nil{
		log.Fatal( e )
	}

  log.Println( ">>> ", t )

 	fmt.Println( gcolage.ReadConfigFile( "test/config.json" ) )

}
