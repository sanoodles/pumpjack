package main

import "fmt"
import "net/http"
import "github.com/kr/pretty"

func main() {
    fmt.Println("hello California")
    
    client := &http.Client{}

    resp, err := client.Get("http://example.com")

    fmt.Printf("%# v", pretty.Formatter(resp));
    fmt.Printf("%# v", pretty.Formatter(err));
}

