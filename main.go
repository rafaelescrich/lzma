package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/klauspost/compress/zlib"
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

	var b bytes.Buffer

	w := zlib.NewWriter(&b)
	w.Write(content)
	w.Close()

	compr := content
	err = ioutil.WriteFile("compressed", compr, 0644)
	if err != nil {
		// Could not obtain stat, handle error
	}

	fmt.Println("content compressed")

	fc, err := os.Stat("compressed")
	if err != nil {
		// Could not obtain stat, handle error
	}
	comprSize := fc.Size()
	fmt.Printf("The compressed file is %d KB long\n", comprSize/1000)

	// TODO: RATIO

}
