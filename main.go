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
	fmt.Println(filenames)

	file, err := os.Open(filenames[0])
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 1000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-c: ", *cFlag)

	fmt.Printf("%d %s\n", count, filenames[0])
}
