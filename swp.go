/*
   swp is a utility for swapping two files.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	file1Name    string
	file2Name    string
	tempFileName string
)

func usage() {
	fmt.Printf("usage: %s [-h] file1 file2\n",
		filepath.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 || os.Args[1] == "-h" {
		usage()
	}

	var err error
	file1Name = os.Args[1]
	file2Name = os.Args[2]
	tempFile, err := ioutil.TempFile("", "swp")
	if err != nil {
		fmt.Println("could not open temporary file: ", err.Error())
		os.Exit(1)
	}
	tempFileName = tempFile.Name()
	err = swap()
	if err != nil {
		fmt.Println("error swapping files: ", err.Error())
	}

	err = os.Remove(tempFileName)
	if err != nil {
		fmt.Println("error removing temporary file: ", err.Error())
	}
}

func swap() error {
	err := writeFile(tempFileName, file1Name)
	if err != nil {
		fmt.Println("error writing temporary file: ", err)
		return err
	}

	err = writeFile(file1Name, file2Name)
	if err != nil {
		fmt.Println("error writing temporary file: ", err)
		return err
	}

	err = writeFile(file2Name, tempFileName)
	if err != nil {
		fmt.Println("error writing temporary file: ", err)
		return err
	}

	return err
}

func writeFile(first, second string) (err error) {
	firstF, err := os.Create(first)
	if err != nil {
		return
	}
	defer firstF.Close()

	secondF, err := os.Open(second)
	if err != nil {
		return err
	}
	defer secondF.Close()

	reader := bufio.NewReader(secondF)
	for {
		var (
			line     []byte
			isPrefix bool
			n        int
		)
		line, isPrefix, err = reader.ReadLine()
		if err == io.EOF {
			err = nil
			break
		} else if isPrefix {
		} else {
			n, err = firstF.Write(line)
			if err != nil {
				return err
			} else if n != len(line) {
				return fmt.Errorf("error writing to line: " +
					err.Error())
			}
			n, err = firstF.WriteString("\n")
			if err != nil {
				return err
			} else if n != 1 {
				return fmt.Errorf("error writing to line: " +
					err.Error())
			}
		}
	}

	return err
}
