package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	readFromFile(args)

}

func readFromFile(args []string) error {

	if len(args) < 2 {
		er := fmt.Errorf("Not enough argument. Provide filename")
		return er
	}

	filename := args[1]
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, f)
	return nil
}
