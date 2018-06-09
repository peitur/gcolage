package gcolage

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func RequestJson(url string) (err error) {
	return nil
}

func RequestFile(url string, dsr string) (err error) {

	target, err := os.Create(dsr)
	if err != nil {
		return err
	}
	defer target.Close()

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected code: %s", res.Status)
	}

	_, err = io.Copy(target, res.Body)
	if err != nil {
		return err
	}

	return nil
}
