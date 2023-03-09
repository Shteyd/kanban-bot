package flag

import (
	"errors"
	"flag"
)

// Returns path to config-file
func Parse() string {
	pathFlag := flag.String("path", "", "path to config-file")

	flag.Parse()

	if *pathFlag == "" {
		panic(errors.New("empty flag error"))
	}

	return *pathFlag
}
