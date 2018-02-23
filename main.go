package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/ulikunitz/xz"
)

func main() {

	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fi, err := os.Stat("file.txt")
	if err != nil {
		// Could not obtain stat, handle error
	}

	origSize := fi.Size()
	fmt.Printf("The original file is %d KB long\n", origSize/1000)

	// const text = "The quick brown fox jumps over the lazy dog.\n"
	var buf bytes.Buffer
	// compress text
	w, err := xz.NewWriter(&buf)
	if err != nil {
		log.Fatalf("xz.NewWriter error %s", err)
	}

	n := bytes.Index(content, []byte{0})
	s := string(content[:n])

	if _, err := io.WriteString(w, s); err != nil {
		log.Fatalf("WriteString error %s", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("w.Close error %s", err)
	}

	err = ioutil.WriteFile("compressed", w., 0644)
	check(err)

	fi, err = os.Stat("compressed")
	if err != nil {
		// Could not obtain stat, handle error
	}

	comprSize := fi.Size()
	fmt.Printf("The compressed file is %d KB long\n", comprSize/1000)

	fmt.Printf("Ratio of compression is %d percent\n", comprSize/origSize*100)

	// // decompress buffer and write output to stdout
	// r, err := xz.NewReader(&buf)
	// if err != nil {
	//     log.Fatalf("NewReader error %s", err)
	// }
	// if _, err = io.Copy(os.Stdout, r); err != nil {
	//     log.Fatalf("io.Copy error %s", err)
	// }
}
