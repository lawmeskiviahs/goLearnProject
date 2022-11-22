package main

import (
    "fmt"
    "log"
    "example.com/greetings"
    "example.com/maths"
    "example/localTest" //modules cannot have underscores in their names or paths
)

func main() {

    log.SetPrefix("greetings: ") // sets prefix to logs
    log.SetFlags(0) // without this date time and line number are also printed in logs
    
    // Get a greeting message and print it.
    message, err := greetings.Hello("Name")
    if err != nil {
        log.Fatal(err) // exits the code with the err
    }
    fmt.Println(message)
	sum := maths.Add(2, 3)
	fmt.Println(sum)
    value:=localTest.Export()
    fmt.Println(value)

}