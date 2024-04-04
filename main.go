package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define commandline flags
	var cFlag = flag.Bool("c", false, "Print out number of bytes in file.")
	flag.Parse()

	filenames := flag.Args()

    if *cFlag {
        for _, filename := range filenames {
            count_bytes(filename)
        }
    }

}

func count_bytes(filename string) {
    data, err := os.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%d %s\n", len(data), filename)
}
