package main

import (
    "flag"
    "log"
    "net"
)

var sock string;

func init() {
    flag.StringVar(&sock, "sock", "/tmp/gorlum.sock", "set unix sock here")
    flag.Parse()
}

func data(c net.Conn) {
    for {
        buf := make([]byte, 512)
        nr, err := c.Read(buf)
        if err != nil {
            return
        }

        data := buf[0:nr]
        println("Server got:", string(data))
        //_, err = c.Write(data)
        if err != nil {
            log.Fatal("Write: ", err)
        }
    }
}

func main() {
    l, err := net.Listen("unix", sock)
    if err != nil {
        log.Fatal("listen error:", err)
    }

    for {
        fd, err := l.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
        }
        go data(fd)
    }
}