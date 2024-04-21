package main

import (
	"fmt"
	"log"
	"net"
)

type (
    FileServer struct {}
)

func (fs *FileServer) start() {
    ln, err := net.Listen("tcp", ":8000")
    if err != nil {
        log.Fatal(err)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go fs.readLoop(conn)
    }
}

func (fs *FileServer) readLoop(conn net.Conn) {
    buf := make([]byte, 2048)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            log.Fatal(err)
        }
        file := buf[:n]
        fmt.Println(file)
        fmt.Printf("received %d bytes over the network\n", n)
    }
}

func main() {
    server := &FileServer{}
    server.start()
}
