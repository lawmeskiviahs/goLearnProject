package main

import (
	// "context"
	"bytes"
	// "encoding/json"
	"fmt"
	"os"
	"log"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {

	// Connect to your local IPFS deamon running in the background.
	content, err := os.ReadFile("bawa.txt")
	// content, err := os.ReadFile("itachi.jpeg")
	if err != nil {
        log.Fatal(err)
    }

	// Where your local node is running on localhost:5001
	india := shell.NewShell("localhost:5001")
	pakistan := shell.NewShell("localhost:5002")

	// Add the file to IPFS
	reader := bytes.NewReader(content)
	cid1, err := india.Add(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid1)

	content2, err := os.ReadFile("bawa2.txt")
	reader2 := bytes.NewReader(content2)
	cid2, err := pakistan.Add(reader2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid2)

	//
	// Get the data from IPFS and output the contents into `struct` format.
	//
	
	
	data, err := india.Cat(cid1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	newStr := buf.String()

	fmt.Println(newStr, "cid1")
buf.Reset()
	// CID2
	dat, err := pakistan.Cat(cid2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	// buff := new(bytes.Buffer)
	buf.ReadFrom(dat)
	newStrr := buf.String()

	fmt.Println(newStrr, "cid2")

	
}