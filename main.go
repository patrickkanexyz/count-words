package main

import (
	"flag"
	"fmt"
	"io"
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

	if len(filenames) == 0 {
		stdin, err := io.ReadAll(os.Stdin)

		if err != nil {
			log.Fatal(err)
		}


        b_count, l_count, w_count, c_count := parse_file(stdin)

        o_string := ""
        if *cFlag {
			o_string += fmt.Sprintf("%d ", b_count)
		}

		if *lFlag {
			o_string += fmt.Sprintf("%d ", l_count)
		}

		if *wFlag {
			o_string += fmt.Sprintf("%d ", w_count)
		}

		if *mFlag {
			o_string += fmt.Sprintf("%d ", c_count)
		}

		if !(*cFlag || *lFlag || *wFlag || *mFlag) {
			fmt.Printf("%d %d %d\n", l_count, w_count, c_count)
		} else {
			fmt.Printf("%s\n", o_string)
		}
        os.Exit(0)
	}

	for _, filename := range filenames {
        data, err := os.ReadFile(filename)
        if err != nil {
            log.Fatal(err)
        }
		byte_count, line_count, word_count, char_count := parse_file(data)

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

		if !(*cFlag || *lFlag || *wFlag || *mFlag) {
			fmt.Printf("%d %d %d %s\n", line_count, word_count, char_count, filename)
		} else {
			fmt.Printf("%s %s\n", output, filename)
		}
	}
}

func parse_file(data []byte) (int, int, int, int) {

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
		IN  = 1
	)

	// Define default state.
	word_state := OUT

	byte_count := len(data)
	line_count := 0
	word_count := 0
	char_count := len(unicode_data)

	for _, runeValue := range unicode_data {
		if unicode.IsSpace(runeValue) {
			//if isWhitespace(runeValue) {
			word_state = OUT

			if runeValue == '\n' {
				line_count++
			}
		}
		//if unicode.IsControl(runeValue) {
		//    fmt.Println("Found control character!")
		//}

		if !unicode.IsSpace(runeValue) {
			//if ! isWhitespace(runeValue) {
			if word_state == OUT {
				word_count++
			}
			word_state = IN
		}
	}

	return byte_count, line_count, word_count, char_count
}
