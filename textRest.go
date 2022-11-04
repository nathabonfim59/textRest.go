package main

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "net/http"
    "os"
    "strconv"
    "flag"
)

// Handles the request
func getRoot(w http.ResponseWriter, r *http.Request) {
    index, err := strconv.Atoi(r.URL.Query().Get("idx"))
    if err != nil {
        fmt.Printf("Invalid index: %v\n", err)
    }

    fmt.Printf("Index: %d \n", index)

    body := readLine(os.Args[1], index)
    io.WriteString(w, body)
}

// Reads a line of a file
func readLine(filename string, lineNum int) (string) {
    f, err := os.Open(filename)
    if err != nil {
        fmt.Printf("An error occurred then opening the file: %v\n", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    var line int

    for scanner.Scan() {
        if line == lineNum {
            return scanner.Text()
        }
        line++
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("An error occurred when reading the file: %v\n", err)
    }

    return ""
}

// Prints the help of the command
func usage() {
    if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" || len(os.Args) < 3 {
        fmt.Printf("usage: %s [filename] [port number]\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(2)
    }
}

func main() {
    usage()

    http.HandleFunc("/", getRoot)
    err := http.ListenAndServe(":" + os.Args[2], nil)

    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("Server closed\n")
    } else if err != nil {
        fmt.Printf("Error starting the server: %s\n", err)
        os.Exit(1)
    }
}

