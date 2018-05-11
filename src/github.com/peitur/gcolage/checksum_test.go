package gcolage

import (
  "testing"
)

func TestChecksumOK( t *testing.T ){
  var checksum = "bfcb0c38bfcbfb64267639d4950c0ebe3fd0117d6c102aa1b9a94379a15df3ab"
  var filename = "test/sample.json"
  var algorithm = "sha256"
  var chs = HashDataFile( filename, algorithm )
  if chs.checksum != checksum{
      t.Error( "TestChecksum = false")
  }
}
