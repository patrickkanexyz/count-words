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
    var mFlag = flag.Bool("m", false, "Print out number of characters in file.")
	flag.Parse()

	filenames := flag.Args()

    for _, filename := range filenames {
        data, err := os.ReadFile(filename)
        if err != nil {
            log.Fatal(err)
        }


        // The data variable is a byte stream. While it mostly behaves like a
        // regular string, it is ASCII formated and not UTF-8/unicode formated.
        // This is a problem because character/rune testing functions from
        // the standard library work on UTF-8/unicode runes and not ASCII.
        //
        // The reason for the data[:] is because we can type cast slices, but
        // not byte streams.
        unicode_data := string(data[:]) 

        // Most of the word count utility can be modeled as a state machine.
        // Here we define the states.
        // OUT: we are outside of a word.
        // IN: we are inside a word.
        const (
            OUT = 0
            IN = 1
        )

        // Define default state.
        word_state := OUT

        byte_count := len(data)
        line_count := 0
        word_count := 0
        char_count := 0

        for _, runeValue := range unicode_data {
            if unicode.IsSpace(runeValue) {
                word_state = OUT

                if runeValue == '\n' {
                    line_count++
                }
            }
            if ! unicode.IsSpace(runeValue) {
                if word_state == OUT {
                    word_count++
                }
                word_state = IN
                char_count++
            }
        }


        // Output string
        output := ""

        if *cFlag {
            output += fmt.Sprintf("%d ", byte_count)
        }

        if *lFlag {
            output += fmt.Sprintf("%d ", line_count)
        }

        if *wFlag {
            output += fmt.Sprintf("%d ", word_count)
        }

        if *mFlag {
            output += fmt.Sprintf("%d ", char_count)
        }

        if ! (*cFlag || *lFlag || *wFlag || *mFlag) {
            fmt.Printf("%d %d %d %s\n", line_count, word_count, char_count, filename)
        } else {
            fmt.Printf("%s %s\n", output, filename)
        }
    }
}
