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
    var lFlag = flag.Bool("l", false, "Print out number of lines in file.")
	flag.Parse()

	filenames := flag.Args()

    if *cFlag {
        for _, filename := range filenames {
            count_bytes(filename)
        }
    }

    if *lFlag {
        for _, filename := range filenames {
            count_lines(filename)
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

func count_lines(filename string) {
    data, err := os.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    count := 0;

    for _, runeValue := range data {
        if runeValue == '\n' {
            count++
        }
    }

    fmt.Printf("%d %s\n", count, filename)
}
