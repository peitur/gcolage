package gcolage

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func RequestJson(url string) ([]byte, error) {
	var res []byte

	buffer, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer buffer.Body.Close()
	if buffer.StatusCode != http.StatusOK {
		return nil, errors.New("Error getting URL")
	}

	fmt.Println(buffer.Body)

	return res, nil
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
