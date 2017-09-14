package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var port = os.Getenv("PORT")

func main() {

	if port == "" {
		port = "5000"
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("error in listening on %s", port)
	}

	fmt.Printf("listening on %s\n", port)
	for {
		if conn, err := ln.Accept(); err == nil {
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {

	var buf bytes.Buffer

	t2conn := io.Writer(&buf)
	tconn := io.TeeReader(conn, t2conn)

	s := strings.Split(strings.Split(buf.String(), "\n")[0], " ")

	var oConn net.Conn
	var err error
	if s[0] == "CONNECT" {
		oConn, err = net.Dial("tcp", s[1])
		conn.Write([]byte(fmt.Sprintf("HTTP/1.1 %d Connection Established\n", http.StatusOK)))
	} else {
		oConn, err = net.Dial("tcp", fmt.Sprintf("%s:80", "www.ishanjain.me"))
	}
	if err != nil {
		conn.Close()
		log.Println(err.Error())
	}
	go io.Copy(conn, oConn)

	io.Copy(oConn, tconn)

}
