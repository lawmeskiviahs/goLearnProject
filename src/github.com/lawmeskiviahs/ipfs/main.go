package main

import (
	"../database"
	// "context"
	// "bytes"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "os"

	// shell "github.com/ipfs/go-ipfs-api"
)

// type Users struct {
	
// 	Cluster string `json:"cluster"`
// }

func main() {

	// var user Users
	// // Connect to your local IPFS deamon running in the background.
	// content, err := os.ReadFile("bawa.txt")
	// // content, err := os.ReadFile("itachi.jpeg")
	// if err != nil {
	// 	log.Fatal(err)
    // }
	
	// jsonn, err := os.Open("data.json")
	// // content, err := os.ReadFile("itachi.jpeg")
	// if err != nil {
	// 	log.Fatal(err)
    // }
	
	// metadata, _ := ioutil.ReadAll(jsonn)
	// json.Unmarshal(metadata, &user)
	// if user.Cluster=="india"{
	// 	// Where your local node is running on localhost:5001
	// 	india := shell.NewShell("localhost:5001")
		
	// 	// Add the file to IPFS
	// 	reader := bytes.NewReader(content)
	// 	cid1, err := india.Add(reader)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "error: %s", err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Printf("added %s\n", cid1)
	// 	data, err := india.Cat(cid1)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "error: %s", err)
	// 		os.Exit(1)
	// 	}
	
	// 	buf := new(bytes.Buffer)
	// 	buf.ReadFrom(data)
	// 	newStr := buf.String()
	
	// 	fmt.Println(newStr, "cid1")
	// 	}	else {
	// 		pakistan := shell.NewShell("localhost:5002")
	// 		content2, err := os.ReadFile("bawa2.txt")
	// 		reader2 := bytes.NewReader(content2)
	// 		cid2, err := pakistan.Add(reader2)
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "error: %s", err)
	// 			os.Exit(1)
	// 		}
	// 		fmt.Printf("added %s\n", cid2)
	// 		dat, err := pakistan.Cat(cid2)
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "error: %s", err)
	// 			os.Exit(1)
	// 		}
		
	// 		// buff := new(bytes.Buffer)
	// 		buf := new(bytes.Buffer)
	// 		buf.ReadFrom(dat)
	// 		newStrr := buf.String()
		
	// 		fmt.Println(ns struct {
	
// 	Cluster string `json:"cluster"`
// }ewStrr, "cid2")
	// 	}
	// 	//
	// 	// Get the data from IPFS and output the contents into `struct` format.
	// 	//
		
		
	// 	// buf.Reset()
	// 	// CID2
	database.Initproc()

	
}