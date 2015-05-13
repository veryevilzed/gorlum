/*
 Демон анализа лог файлов сообщает текущему демону свои наблюдения через unix.sock
*/
package main

import (
    "flag"
    "log"
    "github.com/ActiveState/tail"
)

var file string;

func init() {
    flag.StringVar(&file, "file", "", "set tail file here")
    flag.Parse()
}

func main() {
    //var ip = flag.Int("flagname", 1234, "help message for flagname")
    
    log.Printf("file=%s", file )

    t, _ := tail.TailFile(file, tail.Config{
        Follow: true,
        ReOpen: true})


    for line := range t.Lines {
        log.Printf("%s", line.Text)
    }
}