package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var port = os.Getenv("PORT")

func main() {

	if port == "" {
		port = "8080"
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

	// req, err := http.ReadRequest(reader)

	// if err != nil {
	// 	if err != io.EOF {
	// 		log.Printf("failed to read request: %s", err)
	// 	}
	// 	return
	// }

	// if req.Method != "CONNECT" {
	handleHTTPTraffic(conn)

	// return
	// }

	// b, _ := httputil.DumpRequest(req, true)
	// fmt.Println(string(b))
	//
	// probably trying to connect to a https server and using CONNECT method to do so
	// if req.Method == "CONNECT" {
	//
	// for {
	//
	// con, err := net.Dial("tcp", "localhost:4000")
	// if err != nil {
	// 	log.Printf("error in dialing: %v\n", err)
	// 	return
	// }
	// // conn.Write([]byte(fmt.Sprintf("HTTP/1.1 %d Connection Established\n", http.StatusOK)))

	// go func(conn, con net.Conn) {
	// 	_, err := io.Copy(conn, con)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.Printf("error in reading data: %v", err)
	// 		}
	// 	}
	// }(conn, con)

	// _, err = io.Copy(con, conn)

	// if err != nil {
	// 	if err != io.EOF {
	// 		log.Printf("error in reading data: %v", err)
	// 	}
	// }

	// conn.Close()

	// }

	// } else {

	// most likely http traffic
	// handleHTTPTraffic(conn)

	// }
}

func handleHTTPTraffic(conn net.Conn) {

	for {
		con, err := net.Dial("tcp", "localhost:4000")
		if err != nil {
			log.Printf("error in dialing: %v\n", err)
			return
		}

		defer con.Close()
		// conn.Write([]byte(fmt.Sprintf("HTTP/1.1 %d Connection Established\n", http.StatusOK)))

		go func(conn, con net.Conn) {
			_, err := io.Copy(conn, con)
			if err != nil {
				if err != io.EOF {
					log.Printf("error in reading data: %v", err)
				}
			}
		}(con, conn)

		_, err = io.Copy(conn, con)

		if err != nil {
			if err != io.EOF {
				log.Printf("error in reading data: %v", err)
			}
		}
	}
}
