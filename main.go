package main

import (
	"flag"
	"fmt"
	"log"
	"os"
    "unicode"
)

func main() {
	// Define commandline flags
	var cFlag = flag.Bool("c", false, "Print out number of bytes in file.")
    var lFlag = flag.Bool("l", false, "Print out number of lines in file.")
    var wFlag = flag.Bool("w", false, "Print out number of words in file.")
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

    if *wFlag {
        for _, filename := range filenames {
            count_words(filename)
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

func count_words(filename string) {
    data, err := os.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    const (
        IN = 0
        OUT = 1
        )

    is_word := OUT
    count := 0

    myString := string(data[:])

    for _, runeValue := range myString {
        if unicode.IsSpace(runeValue) {
            is_word = OUT
        }
        if ! unicode.IsSpace(runeValue) {
            if is_word == OUT {
                is_word = IN
                count++
            }
        }
    }

    fmt.Printf("%d %s\n", count, filename)

}

func is_whitespace(c byte) bool {
    if c == '\n' || c == '\t' || c == ' ' {
        return true
    }
    return false
}
