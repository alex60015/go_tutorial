package main

import (
    "fmt"
    "log"
)

import "tutorial/greetings"

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0);

    names := []string{"Alex", "Claudia", "Richard"}

    message, err := greetings.MultipleHellos(names)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}