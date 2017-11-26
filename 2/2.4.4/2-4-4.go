package main

import (
	"io"
	"os"
	"net"
	"net/http"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
//	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	req.Write(conn)
	io.Copy(os.Stdout, conn)
}
