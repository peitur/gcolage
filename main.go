package main

import (
	"fmt"
  "github.com/peitur/gcolage"
)


func main( ){
  fmt.Println("Welcome ...")
  gcolage.PrintHelp()
  var t, _ = gcolage.GetFileInfo( "main.go" )
  fmt.Println( t )
}
