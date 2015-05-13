/*
 Демон анализа лог файлов сообщает текущему демону свои наблюдения через unix.sock
*/
package main

import (
    "flag"
    "net"
    "log"
    "github.com/ActiveState/tail"
)

var file string;
var sock string;

func init() {
    flag.StringVar(&file, "file", "", "set tail file here")
    flag.StringVar(&sock, "sock", "/tmp/gorlum.sock", "set unix sock here")
    flag.Parse()
}

func main() {
    
    log.Printf("file=%s", file )
    c, err := net.Dial("unix", sock)
    if err != nil {
        panic(err)
    }

    t, _ := tail.TailFile(file, tail.Config{
        Follow: true,
        ReOpen: true})


    for line := range t.Lines {
        log.Printf("%s", line.Text)
        _, err := c.Write([]byte(line.Text+"\n"))
        if err != nil {
            log.Fatal("write error:", err)
            break
        }
    }
}