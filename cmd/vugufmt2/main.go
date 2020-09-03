// create flag -w bool
// verify one flag arguement or error message
// take the one arguement read the file in
// if -w write back to the file.
// if no specify print to the console.

package main

import (
	"flag"
	"fmt"
)

func main() {
	writeFlag := flag.Bool("w", false, "-w to wrte to the file")
	flag.Parse()
	_ = writeFlag
	args := flag.Args()
	// if len(args) != 1 {
	// 	log.Fatalf("Please select one")
	// 	openFile, err := os.Open()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	_ = openFile

	fmt.Println(args)
}
