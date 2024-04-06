package basic

import (
	"flag"
	"fmt"
	"os"
)

func Basic() {
	name := flag.String("name", "Shubham", "The name to greet")

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("Hello, %s!\n", *name)
	} else if flag.Arg(0) == "list" {
		files, _ := os.Open(".")
		defer files.Close()

		fileInfo, _ := files.ReadDir(-1)
		for _, file := range fileInfo {
			fmt.Println(file.Name())
		}
	} else {
		fmt.Printf("Hello, %s!\n", *name)
	}
}
