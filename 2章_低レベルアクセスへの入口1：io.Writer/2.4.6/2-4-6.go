package main

import (
	"os"
	"net/http"
)

func main() {

	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダも追加できます")
	request.Write(os.Stdout)
}