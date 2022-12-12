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
	content, err := os.ReadFile("photo.png")
	if err != nil {
        log.Fatal(err)
    }

	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")

	// Add the file to IPFS
	reader := bytes.NewReader(content)
	cid, err := sh.Add(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid)

	//
	// Get the data from IPFS and output the contents into `struct` format.
	//

	data, err := sh.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
}