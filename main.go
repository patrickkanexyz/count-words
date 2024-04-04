package main

import (
	"fmt"
    "flag"
	"os"
    "log"
)

func main() {
    // Define commandline flags
    var cFlag = flag.Bool("c", false, "Print out number of bytes in file.")
    flag.Parse()
    
    filename := os.Args[1]
    fmt.Println(filename)

    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }

    data := make([]byte, 1000)
    count, err := file.Read(data)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("-c: ", *cFlag)

    fmt.Printf("%d %s\n", count, filename)
}
